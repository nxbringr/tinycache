package main

import (
	"tinycache/cache"
	"tinycache/routes"
)

func main() {
	c := cache.NewCache()
	r := routes.SetupRoutes(c)
	r.Run(":8080")
}
