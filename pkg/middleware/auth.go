package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/utility"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		if token != "Bearer token" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		user_id, err := utility.CheckToken(token, config.Config.Server.Secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("user_id", user_id)

		c.Next()
	}
}
