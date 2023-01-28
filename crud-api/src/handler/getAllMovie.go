package handler

import (
	"log"
	"zkfmapf123/crud-api/src/repository"

	"github.com/gin-gonic/gin"
)

var m repository.MovieRepository
func HandleGetAllMovie(c *gin.Context){
	
	movies, err := m.RetrieveAll()
	if err != nil {
		msg := "Movie Retrieve All Error"
		log.Print(msg)
		c.JSON(500, gin.H{
			"message" : msg,
		})
	}

	c.JSON(200, gin.H{
		"movies" : movies,
	})
}