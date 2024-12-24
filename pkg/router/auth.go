package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/pkg/controllers/auth"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
)

func Auth(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *storage.Database, logger *utility.Logger) *gin.Engine {
	authController := auth.Controller{Db: db, Validator: validator, Logger: logger}
	authUrl := r.Group(ApiVersion + "/auth")
	{
		authUrl.POST("/login", authController.LoginUser)
		authUrl.POST("/logout", authController.LogOutUser)

	}
	return r
}
