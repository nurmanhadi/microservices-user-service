package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequest = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "total of number http request",
		},
		[]string{"method", "path", "status"},
	)
	httpLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request latency in seconds",
			Buckets: []float64{.001, .005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"method", "path"},
	)
)

func init() {
	prometheus.MustRegister(httpRequest, httpLatency)
	// prometheus.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	// prometheus.MustRegister(collectors.NewGoCollector())
}
func MetricHttpRequest() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		status := ctx.Writer.Status()
		httpRequest.WithLabelValues(ctx.Request.Method, ctx.FullPath(), fmt.Sprintf("%d", status)).Inc()
	}
}
func MetricHttpLatency() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start).Seconds()
		httpLatency.WithLabelValues(ctx.Request.Method, ctx.FullPath()).Observe(duration)
	}
}
