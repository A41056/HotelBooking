package middlewares

import (
	"github.com/google/uuid"
	"main.go/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is missing"})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
			c.Abort()
			return
		}

		token := authParts[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token is missing"})
			c.Abort()
			return
		}

		userID, err := utils.ValidateJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Convert userID to uuid.UUID
		userUUID, err := uuid.Parse(userID)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid user ID"})
			c.Abort()
			return
		}

		// Get rate limiter for the user
		limiter := getRateLimiter(userUUID)
		if !limiter.limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "too many requests"})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
