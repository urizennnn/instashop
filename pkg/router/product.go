package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/pkg/controllers/products"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
)

func Product(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *storage.Database, logger *utility.Logger) {
	productController := products.Controller{Db: db, Validator: validator, Logger: logger}
	productUrl := r.Group(ApiVersion + "/product")
	{
		productUrl.POST("/create", productController.CreateProduct)
		productUrl.GET("/get", productController.GetProducts)
		productUrl.GET("/get/:id", productController.GetProduct)
		productUrl.PATCH("/update/:id", productController.UpdateProduct)
		productUrl.DELETE("/delete/:id", productController.DeleteProduct)
	}

}
