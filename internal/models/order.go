package models

import (
	"time"

	postgresql "github.com/urizennnn/instashop/pkg/repository/storage/pg"
	"gorm.io/gorm"
)

type Order struct {
	ID          string         `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	UserID      string         `gorm:"column:user_id; type:uuid; not null" json:"user_id"` // Foreign key for User
	ProductID   string         `gorm:"column:product_id; type:uuid; not null" json:"product_id"`
	Quantity    int            `gorm:"column:quantity; type:int" json:"quantity"`
	TotalAmount int            `gorm:"column:total_amount; type:int" json:"total_amount"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`       // Correct relationship
	Product     Product        `gorm:"foreignKey:ProductID" json:"product"` // Add Product relationship
	CreatedAt   time.Time      `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateOrderRequest struct {
	UserID    string `json:"user_id" validate:"required"`
	ProductID string `json:"product_id" validate:"required"`
	Quantity  int    `json:"quantity" validate:"required"`
}

func (o *Order) CreateOrder(db *gorm.DB) error {
	err := postgresql.CreateOneRecord(db, o)
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) UpdateOrder(db *gorm.DB) error {
	var order Order
	err := db.Where("user_id = ?", o.UserID).First(&order).Error
	if err != nil {
		return err
	}
	order.Quantity = o.Quantity
	order.UpdatedAt = time.Now()
	err = db.Save(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) DeleteOrder(db *gorm.DB) error {
	var order Order
	err := db.Where("user_id = ?", o.UserID).First(&order).Error
	if err != nil {
		return err
	}
	err = db.Delete(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) GetOrder(db *gorm.DB) error {
	var order Order
	err := db.Where("user_id = ?", o.UserID).First(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) GetOrderByID(db *gorm.DB, id string) (Order, error) {
	var order Order
	err := db.Where("id = ?", id).First(&order).Error
	if err != nil {
		return order, nil
	}
	return order, nil
}
