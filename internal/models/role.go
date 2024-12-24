package models

import (
	postgresql "github.com/urizennnn/instashop/pkg/repository/storage/pg"
	"gorm.io/gorm"
)

type Role struct {
	ID   string `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	Desc string `gorm:"column:desc; type:varchar(255)" json:"desc"`
	Name string `gorm:"column:name; type:varchar(255)" json:"name"`
}

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required"`
	Desc string `json:"desc" validate:"required"`
}

func (r *Role) CreateRole(db *gorm.DB) error {
	err := postgresql.CreateOneRecord(db, r)
	if err != nil {
		return err
	}
	return nil
}

func (r *Role) UpdateRole(db *gorm.DB) error {
	var role Role
	err := db.Where("name = ?", r.Name).First(&role).Error
	if err != nil {
		return err
	}
	role.Name = r.Name
	err = db.Save(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Role) DeleteRole(db *gorm.DB) error {
	var role Role
	err := db.Where("name = ?", r.Name).First(&role).Error
	if err != nil {
		return err
	}
	err = db.Delete(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Role) FindRoleById(db *gorm.DB, id string) error {
	err := db.Where("id = ?", id).First(&r).Error
	if err != nil {
		return err
	}
	return nil
}
