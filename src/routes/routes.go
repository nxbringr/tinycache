package routes

import (
	"io"
	"net/http"
	"tinycache/cache"
	"tinycache/logger"
	"tinycache/metrics"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *cache.Cache) *gin.Engine {
	metrics.Init()

	r := gin.Default()
	r.Use(metrics.Middleware())
	r.Use(logger.Middleware())

	r.GET("/metrics", metrics.Handler())

	r.GET("/", func(ctx *gin.Context) {
		logger.Log.Infow("Health check hit", "path", ctx.FullPath())
		ctx.JSON(http.StatusOK, "I'm alive!")
	})

	r.POST("/cache/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil || len(body) == 0 {
			logger.Log.Errorw("Invalid POST body", "key", key, "error", err)
			ctx.String(http.StatusBadRequest, "Invalid body")
			return
		}
		c.WriteEntry(key, string(body))
		logger.Log.Infow("Cache entry created", "key", key)
		ctx.Status(http.StatusCreated)
	})

	r.GET("/cache/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		value, ok := c.ReadEntry(key)
		if !ok {
			logger.Log.Warnw("Cache miss", "key", key)
			ctx.String(http.StatusNotFound, "Key not found or expired")
			return
		}
		logger.Log.Infow("Cache hit", "key", key)
		ctx.String(http.StatusOK, value.(string))
	})

	r.DELETE("/cache/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		c.DeleteEntry(key)
		logger.Log.Infow("Cache entry deleted", "key", key)
		ctx.Status(http.StatusNoContent)
	})

	r.NoRoute(func(c *gin.Context) {
		logger.Log.Warnw("Route not found",
			"path", c.Request.URL.Path,
			"method", c.Request.Method,
		)
		c.JSON(http.StatusNotFound, gin.H{
			"error":  "route not found",
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		})
	})

	return r
}
