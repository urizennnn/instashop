package utility

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func ValidatePicture(base64Image string) ([]byte, string, error) {
	const maxImageSize = 5 * 1024 * 1024 // 5MB

	var (
		imageData []byte
		ext       string
	)

	if base64Image == "" {
		return nil, "", nil
	}

	if unsupportedURLPrefix(base64Image) {
		return nil, "", nil
	}

	switch {
	case strings.HasPrefix(base64Image, "data:image/jpeg;base64,"):
		ext = ".jpeg"
	case strings.HasPrefix(base64Image, "data:image/jpg;base64,"):
		ext = ".jpg"
	case strings.HasPrefix(base64Image, "data:image/png;base64,"):
		ext = ".png"
	default:
		return nil, "", fmt.Errorf("invalid content type: only PNG, JPEG, or JPG images are allowed")
	}

	if len(imageData) > maxImageSize {
		return imageData, ext, fmt.Errorf("image size exceeds 5MB limit")
	}

	parts := strings.SplitN(base64Image, ",", 2)

	if len(parts) < 2 {
		return imageData, ext, fmt.Errorf("invalid data URL")
	}
	base64ImageData := parts[1]

	imageData, err := base64.StdEncoding.DecodeString(base64ImageData)
	if err != nil {
		return imageData, ext, fmt.Errorf("failed to decode base64 string: %w", err)
	}

	return imageData, ext, nil
}

func unsupportedURLPrefix(url string) bool {
	supportedPrefixes := []string{"http://", "https://", "blob:", "ipfs://", "ftp:"}
	for _, prefix := range supportedPrefixes {
		if strings.HasPrefix(url, prefix) {
			return true
		}
	}
	return false
}
