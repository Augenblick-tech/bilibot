package email

import (
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Add(email *model.Email) error {
	return dao.Create(email)
}

func GetConfig(UserID uint) (*model.Email, error) {
	email := &model.Email{
		UserID: UserID,
	}
	if err := dao.First(email); err != nil {
		return nil, err
	}
	return email, nil
}
