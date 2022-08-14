package dynamic

import (
	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(dynamics ...bilibot.Dynamic) error {
	// 按最新时间排序，搜索数据库
	old, err := dao.GetDynamicByOder(len(dynamics))
	if err != nil {
		return err
	}

	i, j := checkNew(dynamics, old)

	if i > 0 {
		// dynamics中包含新动态
		return dao.Create(model.ToDynamic(dynamics[:i]...))
	} else if j > 0 {
		// dynamics中只有老动态而且未加入数据库
		return dao.Create(model.ToDynamic(dynamics[len(old)-j:]...))
	}

	// 数据库初始化
	if len(old) == 0 {
		return dao.Create(model.ToDynamic(dynamics...))
	}

	// 动态无更新，无需增加数据
	return nil
}

func Delete(id string) error {
	return dao.Delete(model.Dynamic{DynamicID: id})
}

func GetByMid(mid string) ([]model.Dynamic, error) {
	return dao.GetDynamicByMid(mid)
}

func checkNew(new []bilibot.Dynamic, old []model.Dynamic) (int, int) {
	// 双指针寻找相同的动态
	i, j := 0, 0
	for i < len(new) && j < len(old) {
		if new[i].Modules.Author.PubTS > old[j].PubTS {
			i++
		} else if new[i].Modules.Author.PubTS < old[j].PubTS {
			j++
		} else {
			break
		}
	}

	// 如果数据库中有删除的动态避免导入
	if j > 0 && new[i].Modules.Author.PubTS == old[j].PubTS {
		return i, 0
	}

	return i, j
}
