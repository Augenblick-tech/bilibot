package api

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
)

func Register(c *engine.Context) (interface{}, error) {
	var user = model.User{}

	err := c.Bind(&user)
	if err != nil {
		return nil, err
	}

	err = user.Create()

	return user.Name, err
}

func Login(c *engine.Context) (interface{}, error) {
	var user = model.User{}

	var tempUser = struct {
		Name string `json:"username"`
		Password string `json:"password"`
	}{}

	err := c.Bind(&tempUser)
	if err != nil {
		return nil, err
	}

	err = user.Get(tempUser.Name)
	if err != nil {
		return nil, err
	}

	if user.Password != tempUser.Password {
		return nil, e.RespCode_ParamError
	}

	return user.Name, err
}
