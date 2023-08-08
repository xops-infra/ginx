package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rcrowley/go-metrics"
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

type PathMetrics struct {
	Counters map[string]metrics.Counter
	Timers   map[string]metrics.Timer
}

// Metrics is a middleware function that enables metrics
func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// 记录API请求的处理时间
		duration := time.Since(start)
		metrics.GetOrRegisterTimer("api.latency", nil).Update(duration)

		// 记录API请求的状态码
		statusCode := c.Writer.Status()
		metrics.GetOrRegisterCounter("api.status."+strconv.Itoa(statusCode), nil).Inc(1)
	}
}
