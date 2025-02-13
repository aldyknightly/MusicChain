package models

import (
	"github.com/aldyknightly/MusicChain/api/common"
)

func Migrate() {
	db := common.GetDB()
	db.SetupJoinTable(&User{}, "FavoriteSongs", &UserLikeSong{})                                             //nolint:all
	db.SetupJoinTable(&Playlist{}, "Songs", &PlaylistsSongs{})                                               //nolint:all
	db.AutoMigrate(&User{}, &Song{}, &Comment{}, &Follow{}, &Genre{}, &Playlist{}, &Forget{}, &SongReport{}) //nolint:all

	//create admin account
	// adminEmail, adminPassword := configs.EnvAdmin()
	// hashPassword, _ := common.HashPassword(adminPassword)
	// adminUser := User{Email: adminEmail, Name: "admin", Password: hashPassword, Role: "admin"}
	// common.GetDB().Create(&adminUser)
}
