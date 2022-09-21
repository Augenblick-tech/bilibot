package web

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/services/email"
)

// UpdateSettings godoc
// @Summary     更新设置
// @Description 
// @Tags        web
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       object		body	model.Email	true	"邮件相关设置"
// @Router      /web/setting/update [post]
func UpdateSettings(c *engine.Context) (interface{}, error) {
	UserID := c.Context.GetUint("UserID")
	emailConfig := model.Email{}

	if err := c.Bind(&emailConfig); err != nil {
		return nil, err
	}

	emailConfig.UserID = UserID

	return nil, email.Add(&emailConfig)
}
