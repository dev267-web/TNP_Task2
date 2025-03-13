package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware checks for a valid API key
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey != "secure-key" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
