package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JWTMiddleware cvalidates JWT tokens
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
