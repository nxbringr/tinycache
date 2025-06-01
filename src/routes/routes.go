package routes

import (
	"io"
	"net/http"
	"tinycache/cache"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(c *cache.Cache) *gin.Engine {
	r := gin.Default()

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

	r.GET("/cache", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, c.ReadAllEntries())

	})

	return r

}
