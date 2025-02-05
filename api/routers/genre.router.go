package routers

import (
	"github.com/aldyknightly/MusicChain/api/controllers"
	"github.com/gin-gonic/gin"
)

var genreController = controllers.GenreController{}

func genreRouter(router *gin.RouterGroup) {
	router.GET("/", genreController.GetAll)
}
