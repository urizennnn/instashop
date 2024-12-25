package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/pkg/middleware"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
)

func Setup(logger *utility.Logger, validator *validator.Validate, db *storage.Database, appConfiguration *config.App) *gin.Engine {
	if appConfiguration.Release == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	// Middlewares
	/* r.Use(gin.Logger())
	r.ForwardedByClientIP = true */
	r.Use(middleware.Security())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())

	// routers
	ApiVersion := "api/v1"
	User(r, ApiVersion, validator, db, logger)
	Auth(r, ApiVersion, validator, db, logger)
	Product(r, ApiVersion, validator, db, logger)
	Order(r, ApiVersion, validator, db, logger)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "Welcome to Instashop API",
			"status":  "success",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  "error",
		})
	})

	return r
}
