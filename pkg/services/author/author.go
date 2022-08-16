package author

import (
	"strconv"

	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(mid string, BotID string) error {
	author, err := bilibot.GetInfo(mid)
	if err != nil {
		return err
	}

	return dao.CreateWithIgonreConflict(&model.Author{
		UID:   strconv.Itoa(int(author.Data.Mid)),
		Name:  author.Data.Name,
		Face:  author.Data.Face,
		BotID: BotID,
	})
}

func Del(mid string) error {
	return dao.Delete(&model.Author{UID: mid})
}

func Get(AuthorID string) (*model.Author, error) {
	author := model.Author{
		UID: AuthorID,
	}
	if err := dao.First(&author); err != nil {
		return nil, err
	}
	return &author, nil
}

func GetList(BotID string) ([]*model.Author, error) {
	return dao.GetAuthorList(BotID)
}
