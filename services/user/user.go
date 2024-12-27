package user

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/pkg/repository/storage/redis"
	"github.com/urizennnn/instashop/utility"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
)

func CreateUser(req *models.CreateUserRequest, db *gorm.DB, logger *utility.Logger, ctx *gin.Context, client *storage.Database) (gin.H, int, error) {
	var user models.User
	var Role models.Role
	var (
		email     = req.Email
		firstName = req.FirstName
		lastName  = req.LastName
		password  = req.Password
		role      = req.Role
	)
	err := Role.FindRoleById(db, role)
	if err != nil {
		logger.Error("Error finding role: %v", err)
		return nil, http.StatusBadRequest, err
	}
	password, err = utility.HashPassword(password)
	if err != nil {
		logger.Error("Error hashing password: %v", err)
		return nil, http.StatusBadRequest, err
	}
	user = models.User{
		ID:         uuid.Must(uuid.NewV4()).String(),
		Email:      email,
		FirstName:  firstName,
		LastName:   lastName,
		RoleID:     role,
		Password:   password,
		IsVerified: false,
	}
	otp, err := utility.GenerateOTP(6)
	if err != nil {
		logger.Error("Error generating OTP: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	otpdetails := models.SendOTP{
		OTP:   otp,
		Email: email,
	}

	err = user.CreateUser(db)
	if err != nil {
		logger.Error("Error creating user: %v", err)
		return nil, http.StatusBadRequest, err
	}
	err = SendOtp(&otpdetails, logger)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	expiry := time.Duration(5) * time.Minute
	err = redis.PushtoRedis(redis.Ctx, email, otp, expiry, client.Redis, logger)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	respData := gin.H{
		"message": "User created successfully",
		"user":    user,
		"staus":   http.StatusCreated,
	}
	return respData, http.StatusCreated, nil
}

func SendOtp(req *models.SendOTP, logger *utility.Logger) error {
	htmlContent, err := os.ReadFile("html/otp.html")
	if err != nil {
		logger.Error("Error reading HTML file: %v", err)
		return err
	}

	htmlString := string(htmlContent)
	htmlString = strings.Replace(htmlString, "{{OTP}}", strconv.Itoa(req.OTP), 1)

	m := gomail.NewMessage()
	m.SetHeader("From", config.Config.MAIL.MAIL_SENDER)
	m.SetHeader("To", req.Email)
	m.SetHeader("Subject", "Verify your OTP")
	m.SetBody("text/html", htmlString)

	d := gomail.NewDialer(
		config.Config.MAIL.MAIL_HOST,
		config.Config.MAIL.MAIL_PORT,
		config.Config.MAIL.MAIL_SENDER,
		config.Config.MAIL.MAIL_PASSWORD,
	)

	if err := d.DialAndSend(m); err != nil {
		logger.Error("Error sending email: %v", err)
		return err
	}

	return nil
}

func VerifyOTP(req *models.VerifyOTP, db *gorm.DB, logger *utility.Logger, ctx *gin.Context, client *storage.Database) (gin.H, int, error) {
	var user models.User
	var (
		email = req.Email
		otp   = req.OTP
	)
	value, err := redis.GetfromRedis(redis.Ctx, email, client.Redis, logger)
	if err != nil {
		return nil, http.StatusNotFound, err
	}
	if value == "" {
		return nil, http.StatusNotFound, err
	}
	if value == otp {
		err := user.VerifyUser(db, email)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
	} else {
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "User verified successfully",
		"status":  http.StatusOK,
	}
	return respData, http.StatusOK, nil
}

func ResendOTP(req *models.ResendOTP, ctx *gin.Context, logger *utility.Logger, client *storage.Database) (gin.H, int, error) {
	var (
		email = req.Email
	)
	otp, err := utility.GenerateOTP(6)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}

	otpdetails := models.SendOTP{
		OTP:   otp,
		Email: email,
	}
	err = SendOtp(&otpdetails, logger)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	expiry := time.Duration(5) * time.Minute
	err = redis.PushtoRedis(redis.Ctx, email, otp, expiry, client.Redis, logger)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	respData := gin.H{
		"message": "OTP sent successfully",
		"status":  http.StatusOK,
	}
	return respData, http.StatusOK, nil
}

func UpdateUser(req *models.CreateUserRequest, db *storage.Database, logger *utility.Logger, ctx *gin.Context) (gin.H, int, error) {
	var user models.User
	var (
		email     = req.Email
		firstName = req.FirstName
		lastName  = req.LastName
		password  = req.Password
	)

	if password != "" {
		password, err := utility.HashPassword(password)
		if err != nil {
			return nil, http.StatusBadRequest, err
		}
		user.Password = password
	}

	user.Email = email
	user.FirstName = firstName
	user.LastName = lastName

	err := user.UpdateUser(db.Postgresql)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	respData := gin.H{
		"message": "User updated successfully",
		"user":    user,
		"status":  http.StatusOK,
	}

	return respData, http.StatusOK, nil
}
