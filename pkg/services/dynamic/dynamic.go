package dynamic

import (
	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(dynamics ...bilibot.Dynamic) error {
	// 按最新时间排序，搜索数据库
	// old, err := dao.GetDynamicByOder(len(dynamics))
	// if err != nil {
	// 	return err
	// }

	return dao.Create(model.ToDynamic(dynamics...))
}
