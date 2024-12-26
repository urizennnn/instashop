package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/utility"
	"gorm.io/gorm"
)

func CreateProduct(req *models.CreateProductRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var user models.User
	err := db.Where("id = ?", ctx.GetString("user_id")).First(&user).Error
	if err != nil {
		logger.Error("Failed to get user", err)
		return nil, http.StatusInternalServerError, err
	}
	product := models.Product{
		ID:          utility.GenerateUUID(),
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
		UserID:      ctx.GetString("user_id"),
		User:        user,
	}

	err = db.Create(&product).Error
	if err != nil {
		logger.Error("Failed to create product", err)
		return nil, http.StatusInternalServerError, err
	}

	respData := gin.H{
		"status":  "success",
		"message": "Product created",
		"data":    product,
	}

	return respData, http.StatusCreated, nil
}

func GetProducts(db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var products []models.Product

	err := db.Find(&products).Error
	if err != nil {
		logger.Error("Failed to get products", err)
		return nil, http.StatusInternalServerError, err
	}

	respData := gin.H{
		"status":  "success",
		"message": "Products retrieved",
		"data":    products,
	}

	return respData, http.StatusOK, nil
}

func GetProduct(db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var product models.Product

	err := db.Where("id = ?", ctx.Param("id")).First(&product).Error
	if err != nil {
		logger.Error("Failed to get product", err)
		return nil, http.StatusInternalServerError, err
	}

	respData := gin.H{
		"status":  "success",
		"message": "Product retrieved",
		"data":    product,
	}

	return respData, http.StatusOK, nil
}

func UpdateProduct(req *models.UpdateProductRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var product models.Product

	err := db.Where("id = ?", ctx.Param("id")).First(&product).Error
	if err != nil {
		logger.Error("Failed to get product", err)
		return nil, http.StatusInternalServerError, err
	}


	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Quantity = req.Quantity

	err = db.Save(&product).Error
	if err != nil {
		logger.Error("Failed to update product", err)
		return nil, http.StatusInternalServerError, err
	}

	respData := gin.H{
		"status":  "success",
		"message": "Product updated",
		"data":    product,
	}

	return respData, http.StatusOK, nil
}

func DeleteProduct(db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var product models.Product

	err := db.Where("id = ?", ctx.Param("id")).First(&product).Error
	if err != nil {
		logger.Error("Failed to get product", err)
		return nil, http.StatusNotFound, err
	}

	err = db.Delete(&product).Error
	if err != nil {
		logger.Error("Failed to delete product", err)
		return nil, http.StatusInternalServerError, err
	}

	respData := gin.H{
		"status":  "success",
		"message": "Product deleted",
	}

	return respData, http.StatusOK, nil
}
