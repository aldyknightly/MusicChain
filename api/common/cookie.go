package common

import (
	"net/http"

	"github.com/aldyknightly/MusicChain/api/configs"
	"github.com/gin-gonic/gin"
)

func SetTokenCookie(c *gin.Context, token string) {
	if configs.Environment() == "production" {
		c.SetCookie("access_token", token, 365*60*60*24, "/", "", false, true)
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access_token", token, 365*60*60*24, "/", "", true, true)
}

func ClearTokenCookie(c *gin.Context) {
	if configs.Environment() == "production" {
		c.SetCookie("access_token", "", -1, "/", "", false, true)
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("access_token", "", -1, "/", "", true, true)
}
