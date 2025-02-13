package middlewares

import (
	"fmt"

	"github.com/aldyknightly/MusicChain/api/common"
	"github.com/aldyknightly/MusicChain/api/models"
	"github.com/gin-gonic/gin"
)

type userModel = models.User

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie("access_token")
		if err != nil {
			c.Set("user", nil)
			return
		}

		claims, err := common.DecodeJWT(token.Value)

		if err != nil {
			c.Set("user", nil)
			return
		}

		userID := uint(claims["id"].(float64))
		userCache := common.GetCache(fmt.Sprintf("%v", userID))
		if userCache != nil {
			c.Set("user", userCache)
		} else {
			var user userModel
			err := common.GetDB().First(&user, userID).Error
			if err != nil {
				c.Set("user", nil)
				return
			}
			common.SetCache(fmt.Sprintf("%v", userID), &user)
			c.Set("user", &user)
		}
	}
}
