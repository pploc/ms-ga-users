package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		corrID := c.GetHeader("X-Correlation-ID")
		if corrID == "" {
			corrID = uuid.NewString()
		}
		c.Set("correlationID", corrID)
		c.Header("X-Correlation-ID", corrID)
		c.Next()
	}
}
