package routers

import (
	"github.com/aldyknightly/MusicChain/tree/main/api/controllers"
	"github.com/aldyknightly/MusicChain/tree/main/api/middlewares"
	"github.com/aldyknightly/MusicChain/tree/main/api/middlewares/guard"
	"github.com/gin-gonic/gin"
)

var songController = controllers.SongController{}

func songRouter(router *gin.RouterGroup) {
	router.GET("/", songController.FindMany)
	router.POST("/", guard.Auth(), middlewares.UploadLimit(30), songController.CreateSong)
	router.GET("/:id", songController.GetByID)
	router.POST("/like/:id", guard.Auth(), songController.CreateFavoriteSong)
	router.DELETE("/like/:id", guard.Auth(), songController.DeleteFavoriteSong)
	router.DELETE("/:id", guard.Auth(), guard.Song(), songController.DeleteSong)
	router.PUT("/:id", guard.Auth(), middlewares.UploadLimit(30), guard.Song(), songController.UpdateSong)
	router.POST("/:id/comment", guard.Auth(), songController.CreateComment)
	router.GET("/:id/comment", songController.GetCommentOfSong)
	router.DELETE("/comment/:id", guard.Auth(), guard.Comment(), songController.DeleteComment)
	router.PUT("/comment/:id", guard.Auth(), guard.Comment(), songController.UpdateComment)
	router.POST("/play/:id", songController.IncrementPlayCount)
	router.POST("/:id/report", guard.Auth(), songController.ReportSong)
}
