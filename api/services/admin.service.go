package services

import (
	"github.com/aldyknightly/MusicChain/api/common"
)

type AdminService struct{}

func (AdminService) Ban(userID uint) error {
	return common.GetDB().Model(&userModel{}).Where("id = ?", userID).Update("is_banned", 1).Error
}

func (AdminService) Unban(userID uint) error {
	return common.GetDB().Model(&userModel{}).Where("id = ?", userID).Update("is_banned", 0).Error
}
