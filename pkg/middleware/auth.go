package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/urizennnn/instashop/internal/config"
	"github.com/urizennnn/instashop/internal/models"
	"github.com/urizennnn/instashop/utility"
	"gorm.io/gorm"
)

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

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
func IsAdmin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var role models.Role

		user, err := user.GetUser(db, c.GetString("user_id"))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to find user"})
			c.Abort()
			return
		}
		err = role.FindRoleById(db, user.RoleID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unable to find role"})
			c.Abort()
			return
		}
		println(role.Name)
		if role.Name != "administrator" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}
		c.Next()

	}
}
