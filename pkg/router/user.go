package router

import (
	"github.com/urizennnn/instashop/pkg/controllers/user"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func User(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *storage.Database, logger *utility.Logger) *gin.Engine {
	userController := user.Controller{Db: db, Validator: validator, Logger: logger}
	userUrl := r.Group(ApiVersion + "/user")
	{
		userUrl.POST("/create", userController.CreateUser)
		userUrl.POST("/verify-otp", userController.VerifyOTP)
		userUrl.PATCH("/resend", userController.ResendOTP)
		userUrl.PATCH("/update", userController.UpdateUser)
	}
	return r
}
