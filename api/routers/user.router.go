package routers

import (
	"github.com/aldyknightly/MusicChain/tree/main/api/controllers"
	"github.com/aldyknightly/MusicChain/tree/main/api/middlewares/guard"
	"github.com/gin-gonic/gin"
)

var userController = controllers.UserController{}

func userRouter(router *gin.RouterGroup) {
	router.PUT("/", guard.Auth(), userController.UpdateUser)
	router.GET("/:id", userController.GetUser)
	router.GET("/", userController.SearchUser)
	router.GET("/favorite", guard.Auth(), userController.GetFavoriteSong)
	router.GET("/:id/songs", userController.GetSongOfUser)
	router.GET("/:id/playlists", userController.GetPlaylistOfUser)
	router.POST("/follow/:id", guard.Auth(), userController.Follow)
	router.POST("/unfollow/:id", guard.Auth(), userController.UnFollow)
	router.GET("/follower/:id", userController.GetFollowers)
	router.GET("/following/:id", userController.GetFollowings)
}
