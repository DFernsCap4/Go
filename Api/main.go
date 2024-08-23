package main

import (
	"example/web-service-gin/database"
	"example/web-service-gin/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	if err:= database.Connect(); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	router := gin.Default()
	routes.RegisterAlbumRoutes(router)
	router.Run("localhost:8080")
}