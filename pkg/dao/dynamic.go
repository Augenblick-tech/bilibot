package dao

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func GetDynamic(id uint64) (dynamic model.Dynamic, err error) {
	err = db.DB.Where("id = ?", id).First(&dynamic).Error
	return
}

func GetDynamicByMid(mid string) (dynamic []model.Dynamic, err error) {
	err = db.DB.Where("author_id = ?", mid).Find(&dynamic).Error
	return
}

func GetDynamicByOder(limit int) (dynamic []model.Dynamic, err error) {
	err = db.DB.Limit(limit).Order("dynamic_id desc").Find(&dynamic).Error
	return
}

func GetDynamicList(AuthorID string) ([]*model.Dynamic, error) {
	dynamics := []*model.Dynamic{}
	if err := db.DB.Where("author_id = ?", AuthorID).Find(&dynamics).Error; err != nil {
		return nil, err
	}
	return dynamics, nil
}
