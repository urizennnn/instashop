package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	service "github.com/urizennnn/instashop/services/user"
	"github.com/urizennnn/instashop/utility"
)

type Controller struct {
	Db        *storage.Database
	Validator *validator.Validate
	Logger    *utility.Logger
}

func (C *Controller) CreateUser(ctx *gin.Context) {
	req := models.CreateUserRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		ctx.JSON(http.StatusBadRequest, rd)
		return
	}

	err = C.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusUnprocessableEntity, "error", "Validation failed",
			utility.ValidationResponse(err, C.Validator), nil)
		ctx.JSON(http.StatusUnprocessableEntity, rd)
		return
	}

	resp, status, err := service.CreateUser(&req, C.Db.Postgresql, C.Logger, ctx, C.Db)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to create user", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (C *Controller) VerifyOTP(ctx *gin.Context) {
	var req = models.VerifyOTP{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		ctx.JSON(http.StatusBadRequest, rd)
		return
	}

	err = C.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusUnprocessableEntity, "error", "Validation failed",
			utility.ValidationResponse(err, C.Validator), nil)
		ctx.JSON(http.StatusUnprocessableEntity, rd)
		return
	}

	resp, status, err := service.VerifyOTP(&req, C.Db.Postgresql, C.Logger, ctx, C.Db)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to verify OTP", err, nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (C *Controller) ResendOTP(ctx *gin.Context) {
	var req = models.ResendOTP{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		ctx.JSON(http.StatusBadRequest, rd)
		return
	}

	err = C.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusUnprocessableEntity, "error", "Validation failed",
			utility.ValidationResponse(err, C.Validator), nil)
		ctx.JSON(http.StatusUnprocessableEntity, rd)
		return
	}
	resp, status, err := service.ResendOTP(&req, ctx, C.Logger, C.Db)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to resend OTP", err, nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (C *Controller) UpdateUser(ctx *gin.Context) {
	req := models.CreateUserRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		ctx.JSON(http.StatusBadRequest, rd)
		return
	}

	err = C.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusUnprocessableEntity, "error", "Validation failed",
			utility.ValidationResponse(err, C.Validator), nil)
		ctx.JSON(http.StatusUnprocessableEntity, rd)
		return
	}

	resp, status, err := service.UpdateUser(&req, C.Db, C.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to update user", err, nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

