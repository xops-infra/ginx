package middleware

import (
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
	// 创建一个 PathMetrics 结构体来存储不同路径的计数器和计时器
	pathMetrics := &PathMetrics{
		Counters: make(map[string]metrics.Counter),
		Timers:   make(map[string]metrics.Timer),
	}

	return func(c *gin.Context) {
		// 获取当前请求的路径
		path := c.Request.URL.Path

		// 检查是否已经存在该路径的计数器和计时器，如果不存在则创建新的
		_, ok := pathMetrics.Counters[path]
		if !ok {
			pathMetrics.Counters[path] = metrics.NewCounter()
			metrics.Register(path+"_counter", pathMetrics.Counters[path])
		}
		_, ok = pathMetrics.Timers[path]
		if !ok {
			pathMetrics.Timers[path] = metrics.NewTimer()
			metrics.Register(path+"_timer", pathMetrics.Timers[path])
		}

		// 计数器加一
		pathMetrics.Counters[path].Inc(1)

		// 计时器开始计时
		start := time.Now()

		// 继续处理请求
		c.Next()

		// 计时器停止并记录经过的时间
		elapsed := time.Since(start)
		pathMetrics.Timers[path].Update(elapsed)
	}
}
