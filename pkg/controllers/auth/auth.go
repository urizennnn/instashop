package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	service "github.com/urizennnn/instashop/services/auth"
	"github.com/urizennnn/instashop/utility"
	"net/http"
)

type Controller struct {
	Db        *storage.Database
	Validator *validator.Validate
	Logger    *utility.Logger
}

func (c *Controller) LoginUser(ctx *gin.Context) {
	req := models.LoginRequest{}
	err := ctx.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		ctx.JSON(http.StatusBadRequest, rd)
		return
	}

	err = c.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusUnprocessableEntity, "error", "Validation failed", utility.ValidationResponse(err, c.Validator), nil)
		ctx.JSON(http.StatusUnprocessableEntity, rd)
		return
	}

	resp, status, err := service.LoginUser(&req, c.Db.Postgresql, c.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to login user", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)

}

func (c *Controller) LogOutUser(ctx *gin.Context) {
	resp, status, err := service.LogOutUser(ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to logout user", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}
