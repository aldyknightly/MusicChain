package routers

import (
	"github.com/aldyknightly/MusicChain/api/controllers"
	"github.com/aldyknightly/MusicChain/api/middlewares/guard"
	"github.com/gin-gonic/gin"
)

var authController = controllers.AuthController{}

func authRouter(router *gin.RouterGroup) {
	router.POST("/signup", authController.SignUp)
	router.POST("/signin", authController.SignIn)
	router.POST("/google/signin", authController.GoogleLogin)
	router.POST("/auth", guard.Auth(), authController.Auth)
	router.POST("/logout", authController.LogOut)
	router.POST("/password/forget", authController.ForgetPassword)
	router.POST("/password/reset", authController.ResetPassword)
}
