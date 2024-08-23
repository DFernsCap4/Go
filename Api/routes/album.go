package routes

import(
	"example/web-service-gin/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterAlbumRoutes(router *gin.Engine) {
	router.GET("/albums", handlers.GetAlbums)
	router.GET("/albums/:id", handlers.GetAlbumsById)
	router.POST("/album", handlers.PostAlbums)
}