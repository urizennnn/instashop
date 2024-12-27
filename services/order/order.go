package order

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/utility"
	"gorm.io/gorm"
)

func CreateOrder(req *models.CreateOrderRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var produdct models.Product

	var user models.User
	// Fetch user ID from context
	userID := ctx.GetString("user_id")
	fmt.Println("User ID:", userID)

	// Fetch user from database
	err := db.Where("id = ?", userID).First(&user).Error
	if err != nil {
		logger.Error("Could not get user", err)
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("user with id %s not found", userID)
		}
	}

	// Fetch product from database
	err = db.Where("id = ?", req.ProductID).First(&produdct).Error
	if err != nil {
		logger.Error("Could not get product", err)
		if err == gorm.ErrRecordNotFound {
			return nil, http.StatusNotFound, fmt.Errorf("product with id %s not found", req.ProductID)
		}
		return nil, http.StatusInternalServerError, err
	}

	// Create order
	order := models.Order{
		ProductID:   req.ProductID,
		ID:          utility.GenerateUUID(),
		Quantity:    req.Quantity,
		UserID:      userID,
		TotalAmount: produdct.Price * req.Quantity,
		User:        user,
		Product:     produdct,
		Status:      "pending",
	}

	err = order.CreateOrder(db)
	if err != nil {
		logger.Error("Something went wrong creating your order", err)
		return nil, http.StatusInternalServerError, err
	}

	fmt.Println("Order Created:", order)

	// Return response
	respData := gin.H{
		"message": "Order created successfully",
		"order":   order,
		"status":  http.StatusOK,
	}
	return respData, http.StatusCreated, nil
}

func UpdateOrder(req *models.UpdateOrderRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var order models.Order
	err := db.Where("id = ?", req.OrderID).First(&order).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}
	lower_status := strings.ToLower(req.Status)
	if lower_status != "pending" {
		logger.Error("Cannot update order status")
		return nil, http.StatusBadRequest, err
	}
	order.Status = req.Status
	order.UpdatedAt = time.Now()
	err = db.Save(&order).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "Order updated successfully",
		"status":  http.StatusOK,
	}
	return respData, http.StatusOK, nil
}

func DeleteOrder(id string, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var order models.Order
	err := db.Where("id = ?", id).First(&order).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}
	err = db.Delete(&order).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "Order deleted successfully",
		"status":  http.StatusOK,
	}
	return respData, http.StatusOK, nil
}

func GetOrderByID(id string, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var order models.Order
	err := db.Where("id = ?", id).First(&order).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "Order fetched successfully",
		"order":   order,
		"status":  http.StatusOK,
	}
	return respData, http.StatusOK, nil

}

func GetOrders(db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var orders []models.Order
	err := db.Where("user_id = ?", ctx.GetString("user_id")).Find(&orders).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "Orders fetched successfully",
		"orders":  orders,
		"status":  http.StatusOK,
	}
	return respData, http.StatusOK, nil
}
