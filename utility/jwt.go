package utility

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userId, secret string, expirationTime time.Duration) (string, error) {
	claims := UserClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expirationTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func CheckToken(tokenString, secret string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid or expired token")
	}

	return claims.UserId, nil
}
