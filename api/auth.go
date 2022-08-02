package api

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
)

func Login(c *engine.Context) (interface{}, error) {
	var user = struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	err := c.Bind(&user)

	return user.Username, err
}
