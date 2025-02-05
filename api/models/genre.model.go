package models

import "github.com/aldyknightly/MusicChain/api/common"

type Genre struct {
	ID   uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"unique" json:"name"`
}

func BulkInsertGenre() {
	genres := []Genre{
		{Name: "Pop"},
		{Name: "Hip hop"},
		{Name: "Rock"},
		{Name: "Country"},
		{Name: "Funk"},
		{Name: "Jazz"},
		{Name: "Disco"},
		{Name: "Electronic"},
		{Name: "Blues"},
		{Name: "House"},
		{Name: "Indie"},
		{Name: "R&B"},
		{Name: "Soul"},
		{Name: "Ballad"},
		{Name: "Podcasts"},
		{Name: "Party"},
	}

	common.GetDB().Model(Genre{}).Create(&genres)
}
