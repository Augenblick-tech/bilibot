package web

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/services/email"
)

func UpdateSettings(c *engine.Context) (interface{}, error) {
	UserID := c.Context.GetUint("UserID")
	emailConfig := model.Email{}

	if err := c.Bind(&emailConfig); err != nil {
		return nil, err
	}

	emailConfig.UserID = UserID

	return nil, email.Add(&emailConfig)
}
