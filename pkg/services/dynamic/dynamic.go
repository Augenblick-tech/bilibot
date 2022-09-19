package dynamic

import (
	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(dynamics ...bilibot.Dynamic) error {
	if len(dynamics) == 0 {
		return nil
	}
	return dao.CreateWithIgonreConflict(model.ToDynamic(dynamics...))
}

func Delete(id string) error {
	return dao.Delete(model.Dynamic{DynamicID: id})
}

func GetByMid(mid string, limit int) ([]model.Dynamic, error) {
	return dao.GetDynamicByMid(mid, limit)
}

func GetList(authorID string) ([]*model.Dynamic, error) {
	return dao.GetDynamicList(authorID)
}
