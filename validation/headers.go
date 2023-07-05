package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware for header validation
func ValidateHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for required headers
		requiredHeaders := map[string]string{
			"Content-Type": "Missing required header: Content-Type",
		}

		for header, errorMsg := range requiredHeaders {
			if c.GetHeader(header) == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": errorMsg})
				c.Abort()
				return
			}
		}

		// Validate "Content-Type" header
		contentType := c.GetHeader("Content-Type")
		if contentType != "application/json" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Content-Type header. Expected: application/json"})
			c.Abort()
			return
		}

		c.Next()
	}
}
