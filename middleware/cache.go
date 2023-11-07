package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xops-infra/http-headers"
	"net/http"
	"time"
)

// DisableCache is a middleware function that appends headers
// to prevent the client from caching the HTTP response.
func DisableCache() gin.HandlerFunc {
	return func(c *gin.Context) {
		// hh is the package name alias for http-headers
		// hh.CacheControl is a constant inside hh package
		c.Header(hh.CacheControl, "no-cache, no-store, max-age=0, must-revalidate, value")
		c.Header(hh.Expires, "Thu, 01 Jan 1970 00:00:00 GMT")
		c.Header(hh.LastModified, time.Now().UTC().Format(http.TimeFormat))
		c.Next()
	}
}
