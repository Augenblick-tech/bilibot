package web

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Register(c *engine.Context) (interface{}, error) {
	var user = model.User{}

	err := c.Bind(&user)
	if err != nil {
		return nil, err
	}

	// password encryption

	err = dao.Create(&user)

	return user.Name, err
}

func Login(c *engine.Context) (interface{}, error) {

	var tempUser = struct {
		Name     string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.Bind(&tempUser)
	if err != nil {
		return nil, err
	}
	var user = model.User{
		Name: tempUser.Name,
	}

	err = dao.First(&user)
	if err != nil {
		return nil, err
	}

	// password decryption

	if user.Password != tempUser.Password {
		return nil, e.RespCode_ParamError
	}

	return user.Name, err
}
