package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/patsnapops/http-headers"
	"net/http"
)

const SEP = ", "

// CORS (Cross-Origin Resource Sharing)
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set(hh.AccessControlAllowOrigin, "*")
		c.Writer.Header().Set(hh.AccessControlMaxAge, "86400")
		c.Writer.Header().Set(hh.AccessControlAllowMethods,
			http.MethodConnect+SEP+
				http.MethodDelete+SEP+
				http.MethodGet+SEP+
				http.MethodHead+SEP+
				http.MethodOptions+SEP+
				http.MethodPost+SEP+
				http.MethodPut)
		c.Writer.Header().Set(hh.AccessControlAllowHeaders,
			hh.XRequestedWith+SEP+
				hh.ContentType+SEP+
				hh.Origin+SEP+
				hh.Authorization+SEP+
				hh.Accept+SEP+
				hh.AcceptEncoding+SEP)
		c.Writer.Header().Set(hh.AccessControlExposeHeaders,
			hh.CacheControl+SEP+
				hh.ContentLanguage+SEP+
				hh.ContentLength+SEP+
				hh.ContentType+SEP+
				hh.Expires+SEP+
				hh.LastModified+SEP)
		c.Writer.Header().Set(hh.AccessControlAllowCredentials, "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
