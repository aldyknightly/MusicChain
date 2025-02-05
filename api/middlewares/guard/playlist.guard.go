package guard

import (
	"net/http"

	"github.com/aldyknightly/MusicChain/api/common"
	"github.com/aldyknightly/MusicChain/api/dtos"
	"github.com/aldyknightly/MusicChain/api/models"
	"github.com/gin-gonic/gin"
)

func Playlist() gin.HandlerFunc {
	return func(c *gin.Context) {
		params := dtos.IdParams{}
		if err := c.ShouldBindUri(&params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid playlist ID"})
			return
		}

		user := c.Keys["user"].(*models.User)
		if user.Role == "admin" {
			c.Next()
			return
		}

		var thisPlaylist models.Playlist
		err := common.GetDB().First(&thisPlaylist, params.ID).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "This playlist does not exist"})
			return
		}

		if thisPlaylist.AuthorID != user.ID {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "You are not authorized to perform this action"})
			return
		}
	}
}
