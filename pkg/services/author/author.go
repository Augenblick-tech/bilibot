package author

import (
	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(mid string, BotID uint) error {
	author, err := bilibot.GetInfo(mid)
	if err != nil {
		return err
	}

	return dao.CreateWithIgonreConflict(&model.Author{
		UID:   author.Data.Mid,
		Name:  author.Data.Name,
		Face:  author.Data.Face,
		BotID: BotID,
	})
}

func Del(mid string) error {
	return dao.Delete(&model.Author{UID: mid})
}

func GetList(BotID uint) ([]*model.Author, error) {
	return dao.GetAuthorList(BotID)
}
