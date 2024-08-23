package main

import (
	"example/web-service-gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterAlbumRoutes(router)
	router.Run("localhost:8080")
}