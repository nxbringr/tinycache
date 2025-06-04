package routes

import (
	"io"
	"net/http"
	"tinycache/cache"
	"tinycache/metrics"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *cache.Cache) *gin.Engine {
	metrics.Init()

	r := gin.Default()
	r.Use(metrics.Middleware())
	r.GET("/metrics", metrics.Handler())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "I'm alive!")
	})

	r.POST("/cache/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil || len(body) == 0 {
			ctx.String(http.StatusBadRequest, "Invalid body")
			return
		}
		c.WriteEntry(key, string(body))
		ctx.Status(http.StatusCreated)

	})

	r.GET("/cache/:key", func(ctx *gin.Context) {

		key := ctx.Param("key")
		value, ok := c.ReadEntry(key)
		if !ok {
			ctx.String(http.StatusNotFound, "Key not found or expired")
			return
		}

		ctx.String(http.StatusOK, value.(string))
	})

	r.DELETE("/cache/:key", func(ctx *gin.Context) {
		key := ctx.Param("key")
		c.DeleteEntry(key)
		ctx.Status(http.StatusNoContent)
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"error":  "route not found",
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		})
	})

	return r

}
