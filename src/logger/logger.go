package logger

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

// Init initializes the global logger instance based on ENVIRONMENT (set in .env file)
func Init() {
	var raw *zap.Logger
	var err error

	env := strings.ToLower(os.Getenv("ENVIRONMENT"))
	if env == "production" {
		raw, err = zap.NewProduction()
	} else {
		raw, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	Log = raw.Sugar()
}

// Sync flushes any buffered logs
func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}

// Middleware returns a Gin middleware handler that logs every HTTP request
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		latency := time.Since(start)

		Log.Infow("HTTP request",
			"method", c.Request.Method,
			"path", c.FullPath(),
			"status", c.Writer.Status(),
			"latency_ms", latency.Milliseconds(),
		)

		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				Log.Errorw("Gin error", "error", err.Err)
			}
		}
	}
}
