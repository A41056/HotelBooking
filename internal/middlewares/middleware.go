package middlewares

import (
	"net/http"
	"strings"

	_const "dev.longnt1.git/aessment-hotel-booking.git/internal/const"
	"dev.longnt1.git/aessment-hotel-booking.git/internal/utils"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": _const.ErrAuthorizeHeaderMissing})
			c.Abort()
			return
		}

		authParts := strings.Split(authHeader, " ")
		if len(authParts) != 2 || authParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": _const.ErrInvalidAuthHeaderFormat})
			c.Abort()
			return
		}

		token := authParts[1]
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": _const.ErrMissingToken})
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
			c.JSON(http.StatusUnauthorized, gin.H{"error": _const.ErrInvalidUserID})
			c.Abort()
			return
		}

		// Get rate limiter for the user
		limiter := getRateLimiter(userUUID)
		if !limiter.limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": _const.ErrTooManyRequest})
			c.Abort()
			return
		}

		c.Set("userID", userID)

		c.Next()
	}
}
