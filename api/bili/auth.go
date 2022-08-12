package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
)

// GetLoginUrl godoc
// @Summary      获取二维码登录链接
// @Description
// @Tags         v2
// @Produce      json
// @Router       /bili/qrcode/getLoginUrl [get]
func GetLoginUrl(c *engine.Context) (interface{}, error) {
	qrcode, err := bilibot.GetLoginUrl()
	if err != nil {
		return nil, err
	}

	return qrcode, nil
}

// GetLoginInfo godoc
// @Summary      获取二维码状态
// @Description
// @Tags         v2
// @Accept       json
// @Produce      json
// @Param        qrcode   body     string  true  "oauthKey"
// @Router       /bili/qrcode/getLoginInfo [post]
func GetLoginInfo(c *engine.Context) (interface{}, error) {
	var oauth = struct {
		OauthKey string `json:"oauthKey" binding:"required"`
	}{}

	err := c.Bind(&oauth)
	if err != nil {
		return nil, err
	}

	cookie, err := bilibot.GetLoginInfo(oauth.OauthKey)
	if err != nil {
		return nil, err
	}

	return cookie, nil
}
