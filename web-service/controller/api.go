package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"web-service-gin/model"
)

var Albums []model.Album

// GetAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Albums)
}

func PostAlbums(c *gin.Context) {
	var newAlbum model.Album
	// Call BindJSON to bind the received JSON to
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	log.Println(fmt.Sprintf("Adding %s to Album", newAlbum.ID))
	// Add the new album to the slice.
	Albums = append(Albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
