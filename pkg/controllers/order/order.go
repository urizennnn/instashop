package order

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	service "github.com/urizennnn/instashop/services/order"
	"github.com/urizennnn/instashop/utility"
	"net/http"
)

type Controller struct {
	Db        *storage.Database
	Logger    *utility.Logger
	Validator *validator.Validate
}

func (c *Controller) CreateOrder(ctx *gin.Context) {
	req := models.CreateOrderRequest{}
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

	resp, status, err := service.CreateOrder(&req, c.Db.Postgresql, c.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to create order", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (c *Controller) GetOrders(ctx *gin.Context) {
	resp, status, err := service.GetOrders(c.Db.Postgresql, c.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to get orders", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (c *Controller) GetOrder(ctx *gin.Context) {
	var order_id = ctx.Param("id")
	resp, status, err := service.GetOrderByID(order_id, c.Db.Postgresql, c.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to get order", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (c *Controller) UpdateOrder(ctx *gin.Context) {
	req := models.UpdateOrderRequest{}
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

	resp, status, err := service.UpdateOrder(&req, c.Db.Postgresql, c.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to update order", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}

func (c *Controller) DeleteOrder(ctx *gin.Context) {
	var order_id = ctx.Param("id")
	resp, status, err := service.DeleteOrder(order_id, c.Db.Postgresql, c.Logger, ctx)
	if err != nil {
		rd := utility.BuildErrorResponse(status, "error", "Failed to delete order", err.Error(), nil)
		ctx.JSON(status, rd)
		return
	}

	ctx.JSON(status, resp)
}
