package middlewares

import (
	"strings"

	"github.com/Lan-shi/task-5-pbi-btpns--Ahmad-_Maulana_Indidharmanto-/services"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := services.ValidateToken(token)
		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
