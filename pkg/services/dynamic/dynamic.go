package dynamic

import (
	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(dynamics ...bilibot.Dynamic) error {
	return dao.CreateWithIgonreConflict(model.ToDynamic(dynamics...))
}

func Delete(id string) error {
	return dao.Delete(model.Dynamic{DynamicID: id})
}

func GetByMid(mid string) ([]model.Dynamic, error) {
	return dao.GetDynamicByMid(mid)
}

func GetList(authorID string) ([]*model.Dynamic, error) {
	return dao.GetDynamicList(authorID)
}
