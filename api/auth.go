package api

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/e"
)

func Login(c *engine.Context) (interface{}, error) {
	user := c.PostBody()
	if len(user) == 0 {
		return nil, e.RespCode_ParamError
	}

	username := user["username"].(string)
	password := user["password"].(string)

	if username == "" || password == "" {
		return nil, e.RespCode_ParamError
	}

	return len(user), nil
}
