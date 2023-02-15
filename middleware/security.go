package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/patsnapops/http-headers"
)

// Secure is a middleware function that appends security and resource access headers
func Secure() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header(hh.XContentTypeOptions, "nosniff")
		c.Header(hh.XFrameOptions, "DENY")
		c.Header(hh.XXSSProtection, "1; mode=block")
		if c.Request.TLS != nil {
			c.Header(hh.StrictTransportSecurity, "max-age=31536000")
		}

		// Also consider adding Content-Security-Policy headers
		// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
	}
}
