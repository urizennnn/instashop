package utility

import (
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
func CheckToken() {}
