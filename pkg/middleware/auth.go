package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/pkg/repository/storage"
	"github.com/urizennnn/instashop/utility"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		println(token)
		if !strings.HasPrefix(token, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(token, " ")
		if len(tokenParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		actualToken := tokenParts[1]

		userID, err := utility.CheckToken(actualToken, config.Config.Server.Secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Expired or invalid token"})
			c.Abort()
			return
		}

		c.Set("user_id", userID)

		c.Next()
	}
}
func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var role models.Role

		user, err := user.GetUser(*&storage.DB.Postgresql, c.GetString("user_id"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		err = role.FindRoleById(*&storage.DB.Postgresql, user.RoleID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		if role.Name != "administator" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()

	}
}
