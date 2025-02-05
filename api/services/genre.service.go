package services

import (
	"github.com/aldyknightly/MusicChain/tree/main/api/common"
	"github.com/aldyknightly/MusicChain/tree/main/api/models"
)

type GenreService struct{}

func (GenreService) GetAll() ([]models.Genre, error) {
	var genres []models.Genre
	err := common.GetDB().Order("id asc").Find(&genres).Error

	return genres, err
}
