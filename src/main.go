package main

import (
	"tinycache/cache"
	"tinycache/logger"
	"tinycache/routes"
)

func main() {
	logger.Init()
	defer logger.Sync()

	c := cache.NewCache()
	r := routes.SetupRoutes(c)
	r.Run(":8080")
}
