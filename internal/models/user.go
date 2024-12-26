package models

import (
	"time"

	postgresql "github.com/urizennnn/instashop/pkg/repository/storage/pg"
	"gorm.io/gorm"
)

type User struct {
	ID         string         `gorm:"type:uuid;primaryKey;unique;not null" json:"id"`
	FirstName  string         `gorm:"column:first_name; type:varchar(255)" json:"first_name"`
	LastName   string         `gorm:"column:last_name; type:varchar(255)" json:"last_name"`
	Email      string         `gorm:"column:email;unique;not null; type:varchar(255)" json:"email"`
	IsVerified bool           `gorm:"column:is_verified; type:bool" json:"is_verified"`
	RoleID     string         `gorm:"type:uuid;not null" json:"role_id"`
	Product    []Product      `gorm:"foreignKey:UserID" json:"product"`
	Order      []Order        `gorm:"foreignKey:UserID" json:"order"`
	Password   string         `gorm:"column:password; type:text; not null" json:"-"`
	CreatedAt  time.Time      `gorm:"column:created_at; not null; autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at; null; autoUpdateTime" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
type VerifyOTP struct {
	OTP   string `json:"otp" validate:"required"`
	Email string `json:"email" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type SendOTP struct {
	OTP   int
	Email string
}
type ResendOTP struct {
	Email string `json:"email" validate:"required"`
}

func (u *User) CreateUser(db *gorm.DB) error {
	err := postgresql.CreateOneRecord(db, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(db *gorm.DB) error {
	var user User
	err := db.Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		return err
	}
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Password = u.Password
	user.IsVerified = u.IsVerified
	user.Email = u.Email
	user.UpdatedAt = time.Now()
	err = db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) VerifyUser(db *gorm.DB, email string) error {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return err
	}
	user.IsVerified = true
	err = db.Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}
func (u *User) LoginUser(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func (u *User) GetUser(db *gorm.DB, id string) (User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
