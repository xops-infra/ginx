package middleware

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func AttachGzipCompress(ginEngine *gin.Engine) *gin.Engine {
	ginEngine.Use(gzip.Gzip(gzip.DefaultCompression))
	return ginEngine
}
