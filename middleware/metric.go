package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	apiCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "pop_api_metrics",
			Help: "The total number of processed events",
		},
		[]string{"method", "api_path"},
	)
	reqDuration = promauto.NewHistogram(prometheus.HistogramOpts{
		Name: "request_duration_seconds",
		Help: "Time taken fulfilling requests",
		ConstLabels: map[string]string{
			"service": "pop",
		},
	})
)

type MetricInfluxConfig struct {
	Host        string
	Database    string
	Measurement string
	Username    string
	Password    string
}

func NewMetricInfluxConfig(host string, database string, measurement string, username string, password string) *MetricInfluxConfig {
	return &MetricInfluxConfig{
		Host:        host, // 127.0.0.1:8086
		Database:    database,
		Measurement: measurement,
		Username:    username,
		Password:    password,
	}
}

// Metrics is a middleware function that enables metrics
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		// 记录API请求的处理时间
		duration := time.Since(start)
		reqDuration.Observe(duration.Seconds())
		// 记录API请求的状态码 加入到api.status的counter中 不能加入到 registry中，而是加入到 bucket中
		apiCounter.With(prometheus.Labels{"api_path": c.Request.URL.Path, "method": c.Request.Method}).Inc()
	}
}
