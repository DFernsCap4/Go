package handlers

import (
	"net/http"
	"example/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func GetAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsById(c *gin.Context){
	id := c.Param("id")

	for _,value := range albums{
		if value.ID == id {
			c.IndentedJSON(http.StatusOK, value)
		}
	}

}

func PostAlbums(c *gin.Context){
	var newAlbum models.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
}