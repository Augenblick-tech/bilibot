package api

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/db"
)

func Register(c *engine.Context) (interface{}, error) {
	var user = db.User{}

	err := c.Bind(&user)
	if err != nil {
		return nil, err
	}

	// password encryption

	err = user.Create()

	return user.Name, err
}

func Login(c *engine.Context) (interface{}, error) {
	var user = db.User{}

	var tempUser = struct {
		Name     string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.Bind(&tempUser)
	if err != nil {
		return nil, err
	}

	err = user.Find(tempUser.Name)
	if err != nil {
		return nil, err
	}

	// password decryption

	if user.Password != tempUser.Password {
		return nil, e.RespCode_ParamError
	}

	return user.Name, err
}
