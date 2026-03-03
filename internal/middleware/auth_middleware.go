package middleware

import (
	"net/http"

	"ms-ga-user/pkg/utils"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		if userID == "" {
			utils.ErrResponse(c, http.StatusUnauthorized, "missing X-User-ID header")
			c.Abort()
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}

func RequirePermission(permission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		perms := c.GetHeader("X-User-Permissions")
		// if permission not in perms
		// utils.ErrResponse(c, http.StatusForbidden, "insufficient permissions")
		// c.Abort()
		_ = perms // mock bypass for now
		c.Next()
	}
}
