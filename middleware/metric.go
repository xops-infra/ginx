package middleware

import (
	"sync"
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

var (
	statusCodeMetrics = make(map[int]metrics.Counter)
	lock              sync.Mutex
)

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()

		// 记录API请求的处理时间
		duration := time.Since(start)
		metrics.GetOrRegisterTimer("api.latency", nil).Update(duration)

		// 记录API请求的状态码
		statusCode := c.Writer.Status()
		registerStatusCodeMetric(statusCode).Inc(1)
	}
}

func registerStatusCodeMetric(code int) metrics.Counter {
	lock.Lock()
	defer lock.Unlock()

	metric, ok := statusCodeMetrics[code]
	if !ok {
		metric = metrics.NewCounter()
		metrics.GetOrRegister("api.status", metric)
		statusCodeMetrics[code] = metric
	}

	return metric
}
