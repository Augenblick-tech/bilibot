package dao

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func GetUserByName(username string) (user model.User, err error) {
	err = db.DB.Where("name = ?", username).First(&user).Error
	return
}

func GetUserByID(id string) (user model.User, err error) {
	err = db.DB.Where("id = ?", id).First(&user).Error
	return
}
