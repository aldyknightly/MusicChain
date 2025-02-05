package helper

import (
	"github.com/aldyknightly/MusicChain/api/models"
	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) uint {
	user := c.Keys["user"]
	if user == nil {
		return 0
	}
	return user.(*models.User).ID
}
