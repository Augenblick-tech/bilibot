package bili

import (
	"net/http"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
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

	return cookie, bot.Add(cookie, 1) // UserID 暂时设为 1
}

// CheckLogin godoc
// @Summary      查询登陆状态
// @Description
// @Tags         v2
// @Accept       json
// @Produce      json
// @Param        SESSDATA   body     string  true  "SESSDATA"
// @Router       /bili/qrcode/check [post]
func CheckLogin(c *engine.Context) (interface{}, error) {
	var cookie = struct {
		SESSDATA string `json:"SESSDATA" binding:"required"`
	}{}

	err := c.Bind(&cookie)
	if err != nil {
		return nil, err
	}
	return bilibot.GetBotInfo(&http.Cookie{Name: "SESSDATA", Value: cookie.SESSDATA})
}
