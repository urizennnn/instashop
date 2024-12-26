package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/pkg/controllers/order"
	"github.com/urizennnn/instashop/pkg/middleware"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
)

func Order(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *storage.Database, logger *utility.Logger) {
	orderController := order.Controller{Db: db, Validator: validator, Logger: logger}
	orderUrl := r.Group(ApiVersion+"/order", middleware.ValidateToken(), middleware.IsAdmin(db.Postgresql))
	{
		orderUrl.POST("/create", orderController.CreateOrder)
		orderUrl.GET("/get", orderController.GetOrders)
		orderUrl.GET("/get/:id", orderController.GetOrder)
		orderUrl.PUT("/update", orderController.UpdateOrder)
		orderUrl.DELETE("/delete/:id", orderController.DeleteOrder)
	}
}
