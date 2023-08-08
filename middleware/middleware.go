package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
	influxdb "github.com/vrischmann/go-metrics-influxdb"
)

type Middleware struct {
	ginEngine *gin.Engine
}

func AttachTo(ginEngine *gin.Engine) *Middleware {
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

// WithRecover is a middleware function that recovers from any panics and writes a 500 if there was one
func (m *Middleware) WithRecover() *Middleware {
	m.ginEngine.Use(Recover())
	return m
}

// WithRequestID is a middleware function that sets a request ID
func (m *Middleware) WithRequestID(key string) *Middleware {
	m.ginEngine.Use(RequestID(key))
	return m
}

// WithSecurity is a middleware function that enhance security
func (m *Middleware) WithSecurity() *Middleware {
	m.ginEngine.Use(Secure())
	return m
}

// WithMetrics is a middleware function that enables metrics
func (m *Middleware) WithMetrics(config *MetricInfluxConfig) *Middleware {
	go influxdb.InfluxDB(metrics.DefaultRegistry,
		5*time.Second, // 1 seconds
		config.Host,
		config.Database,
		config.Measurement,
		config.Username,
		config.Password,
		true, // enable aligned timestamps
	)
	go metrics.Log(metrics.DefaultRegistry, 5*time.Second, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
	m.ginEngine.Use(Metrics())
	return m
}
