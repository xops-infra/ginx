package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	ginEngine *gin.Engine
}

func New(ginEngine *gin.Engine) *Middleware {
	return &Middleware{
		ginEngine: ginEngine,
	}
}

// WithCacheDisabled is a middleware function that sets cache headers
func (m *Middleware) WithCacheDisabled() *Middleware {
	m.ginEngine.Use(DisableCache())
	return m
}

// WithCompress is a middleware function that compresses the response
func (m *Middleware) WithCompress() *Middleware {
	m.ginEngine = AttachGzipCompress(m.ginEngine)
	return m
}

// WithCORS is a middleware function that allow CORS
func (m *Middleware) WithCORS() *Middleware {
	m.ginEngine.Use(CORS())
	return m
}

// WithRequestID is a middleware function that sets a request ID
func (m *Middleware) WithRequestID(key string) *Middleware {
	m.ginEngine.Use(RequestID(key))
	return m
}

// WithRecover is a middleware function that recovers from any panics and writes a 500 if there was one
func (m *Middleware) WithRecover() *Middleware {
	m.ginEngine.Use(Recover())
	return m
}

// WithSecurity is a middleware function that enhance security
func (m *Middleware) WithSecurity() *Middleware {
	m.ginEngine.Use(Secure())
	return m
}
