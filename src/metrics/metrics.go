package metrics

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var httpRequestsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests labeled by path, method, status.",
	},
	[]string{"path", "method", "status"},
)

func Init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		status := fmt.Sprintf("%d", c.Writer.Status())
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		httpRequestsTotal.WithLabelValues(path, c.Request.Method, status).Inc()
	}
}

func Handler() gin.HandlerFunc {
	return gin.WrapH(promhttp.Handler())
}
