package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID(key string) gin.HandlerFunc {
	// normally the key is X-Request-ID
	return func(c *gin.Context) {
		id := uuid.New().String()
		c.Set(key, id)
		c.Writer.Header().Set(key, id)
		c.Next()
	}
}
