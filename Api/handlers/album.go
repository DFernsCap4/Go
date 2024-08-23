package handlers

import (
	"example/web-service-gin/database"
	"example/web-service-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetAlbums(c *gin.Context){
	var albums []models.Album
	result := database.DB.Find(&albums)
	
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbumsById(c *gin.Context){
	id := c.Param("id")

	var album models.Album

	result := database.DB.First(&album, id)


	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		}else{
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func PostAlbums(c *gin.Context){
	var newAlbum models.Album

	if err:= c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	result := database.DB.Create(&newAlbum)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAlbum)
}