package models

import (
	postgresql "github.com/urizennnn/instashop/pkg/repository/storage/pg"
	"gorm.io/gorm"

	"time"
)

type Product struct {
	ID          string         `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	Name        string         `gorm:"column:name; type:varchar(255)" json:"name"`
	Description string         `gorm:"column:description; type:text" json:"description"`
	Price       int            `gorm:"column:price; type:int" json:"price"`
	Quantity    int            `gorm:"column:quantity; type:int" json:"quantity"`
	UserID      string         `gorm:"column:user_id; type:uuid; not null" json:"user_id"`
	User        User           `gorm:"foreignKey:UserID" json:"user"`
	CreatedAt   time.Time      `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
type CreateProductRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

type UpdateProductRequest struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
	Quantity    int    `json:"quantity" validate:"required"`
}

func (p *Product) CreateProduct(db *gorm.DB) error {
	err := postgresql.CreateOneRecord(db, p)
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) UpdateProduct(db *gorm.DB) error {
	var product Product
	err := db.Where("name = ?", p.Name).First(&product).Error
	if err != nil {
		return err
	}
	product.Name = p.Name
	product.Description = p.Description
	product.Price = p.Price
	product.Quantity = p.Quantity
	product.UpdatedAt = time.Now()
	err = db.Save(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Product) DeleteProduct(db *gorm.DB) error {
	var product Product
	err := db.Where("name = ?", p.Name).First(&product).Error
	if err != nil {
		return err
	}
	err = db.Delete(&product).Error
	if err != nil {
		return err
	}
	return nil
}

func GetProduct(db *gorm.DB) ([]Product, error) {
	var product []Product
	err := db.Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func GetProductByID(db *gorm.DB, id string) (Product, error) {
	var product Product
	err := db.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
