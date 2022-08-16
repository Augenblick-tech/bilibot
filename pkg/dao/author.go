package dao

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func GetAuthorList(BotID string) ([]*model.Author, error) {
	authors := make([]*model.Author, 0)
	if err := db.DB.Find(&authors, "bot_id = ?", BotID).Error; err != nil {
		return nil, err
	}
	return authors, nil
}
