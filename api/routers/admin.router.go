package routers

import (
	"github.com/aldyknightly/MusicChain/tree/main/api/controllers"
	"github.com/aldyknightly/MusicChain/tree/main/api/middlewares/guard"
	"github.com/gin-gonic/gin"
)

var adminController = controllers.AdminController{}

func adminRouter(router *gin.RouterGroup) {
	router.POST("/ban/:id", guard.Auth(), guard.AdminGuard(), adminController.BanUser)
	router.DELETE("/ban/:id", guard.Auth(), guard.AdminGuard(), adminController.UnbanUser)
}
