package dao

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func GetBotList(UserID uint) ([]*model.Bot, error) {
	bots := []*model.Bot{}
	if err := db.DB.Where("user_id = ?", UserID).Find(&bots).Error; err != nil {
		return nil, err
	}
	return bots, nil
}