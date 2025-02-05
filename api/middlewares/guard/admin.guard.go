package guard

import (
	"net/http"

	"github.com/aldyknightly/MusicChain/api/models"
	"github.com/gin-gonic/gin"
)

func AdminGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.Keys["user"].(*models.User)
		if user.Role != "admin" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Only admin can perform this action"})
			return
		}
	}
}
