package utility

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type UserClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(db *gorm.DB, userId, secret string) (string, error) {
	claims := UserClaims{
		UserId:         userId,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))

}
func CheckToken(tokenString, secret string) (string, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	// Extract claims and verify
	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	return claims.UserId, nil
}
