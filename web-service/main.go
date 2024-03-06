package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"web-service-gin/controller"
	"web-service-gin/model"
)

func main() {
	log.Println("Executing Go Tasks")
	controller.Albums = []model.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}
	// Album struct must be of exported type
	router := gin.Default()
	router.GET("/albums", controller.GetAlbums)
	router.POST("/add/album", controller.PostAlbums)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}
