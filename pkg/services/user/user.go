package user

import (
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(user model.User) error {
	return dao.Create(&user)
}

func Get(username string) (model.User, error) {
	return dao.GetUserByName(username)
}
