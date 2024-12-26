package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/utility"
	"gorm.io/gorm"
)

func LoginUser(req *models.LoginRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (interface{}, int, error) {
	var user models.User

	err := db.Where("email = ?", req.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("user not found")
		}
		return nil, http.StatusInternalServerError, err
	}

	if !utility.CompareHash(req.Password, user.Password) {
		return nil, http.StatusUnauthorized, errors.New("invalid password")
	}

	if !user.IsVerified {
		return nil, http.StatusUnauthorized, errors.New("user not verified")
	}

	token, err := utility.GenerateToken(user.ID, config.Config.Server.Secret, time.Hour*5)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	respData := gin.H{
		"token":   token,
		"status":  "success",
		"message": "login success",
	}

	return respData, http.StatusOK, nil
}

func LogOutUser(ctx *gin.Context) (interface{}, int, error) {
	respData := gin.H{
		"status":  "success",
		"message": "logout success",
	}

	return respData, http.StatusOK, nil
}
