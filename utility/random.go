package utility

import (
	crand "crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/gofrs/uuid"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func GetRandomNumbersInRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func RandomString(length int) string {
	u, _ := uuid.NewV4()
	uuidStr := u.String()
	// Regular expression pattern to match all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return ""
	}
	// Replacing all non-alphanumeric characters with an empty string
	processedString := reg.ReplaceAllString(uuidStr+uuidStr[:length%36], "")
	if len(processedString) >= length {
		return processedString[:length]
	}
	// Padding the processed string with random alphanumeric characters to make it the desired length
	alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	padding := make([]byte, length-len(processedString))
	rand.Read(padding)
	for i, b := range padding {
		padding[i] = alphanumeric[b%byte(len(alphanumeric))]
	}
	return processedString + string(padding)
}

func GenerateOTP(max int) (int, error) {
	b := make([]byte, max)
	n, err := io.ReadFull(crand.Reader, b)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return strconv.Atoi(string(b))
}

func GenerateInvitationToken() (string, error) {
	bytes := make([]byte, 16)
	_, err := crand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateInvitationLink(baseurl, orgID, token string) string {
	return baseurl + fmt.Sprintf("/accept_org_invitation?org_id=%s&invitation_token=%s", orgID, token)
}
func GenerateChannelInvitationLink(baseurl, channelID, token string) string {
	return baseurl + fmt.Sprintf("/accept_channel_invitation?channel_id=%s&invitation_token=%s", channelID, token)
}
