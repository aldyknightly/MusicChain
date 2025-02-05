package controllers

import (
	"net/http"

	"github.com/aldyknightly/MusicChain/tree/main/api/services"
	"github.com/gin-gonic/gin"
)

type GenreController struct{}

var genreService = services.GenreService{}

func (GenreController) GetAll(c *gin.Context) {
	genres, err := genreService.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, &genres)
}
