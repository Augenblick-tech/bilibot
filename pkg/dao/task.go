package dao

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func GetAllTask() ([]*model.Task, error) {
	tasks := make([]*model.Task, 0)
	if err := db.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
