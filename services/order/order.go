package order

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/utility"
	"gorm.io/gorm"
)

func CreateOrder(req *models.CreateOrderRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var user models.User
	var produdct models.Product

	user, err := user.LoginUser(db, ctx.GetString("user_id"))
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}
	err = db.Where("id = ?", req.ProductID).First(&produdct).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}
	order := models.Order{
		ProductID:   req.ProductID,
		Quantity:    req.Quantity,
		UserID:      ctx.GetString("user_id"),
		TotalAmount: produdct.Price * req.Quantity,
	}

	err = order.CreateOrder(db)
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "Order created successfully",
		"order":   order,
		"status":  http.StatusOK,
	}
	return respData, http.StatusCreated, nil

}

func UpdateOrder(req *models.UpdateOrderRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var order models.Order
	err := db.Where("user_id = ?", ctx.GetString("user_id")).First(&order).Error
	if err != nil {
		logger.Error(err)
		return nil, http.StatusBadRequest, err
	}
	order.Quantity = req.Quantity
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

func DeleteOrder(db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var order models.Order
	err := db.Where("user_id = ?", ctx.GetString("user_id")).First(&order).Error
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

func GetOrder(db *gorm.DB, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var order models.Order
	err := db.Where("user_id = ?", ctx.GetString("user_id")).First(&order).Error
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
	err := db.Find(&orders).Error
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
