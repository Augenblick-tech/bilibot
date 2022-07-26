package bili

import (
	"net/http"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model/api"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
)

// GetLoginUrl godoc
// @Summary     获取二维码登录链接
// @Description
// @Tags        bili
// @Produce     json
// @Security 	ApiKeyAuth
// @Success		200			{object}	api.BiliQrCodeAuth
// @Router      /bili/qrcode/getLoginUrl [get]
func GetLoginUrl(c *engine.Context) (interface{}, error) {
	qrcode, err := bilibot.GetLoginUrl()
	if err != nil {
		return nil, err
	}

	return api.BiliQrCodeAuth{
		TS:        qrcode.TS,
		Url:       qrcode.Data.Url,
		QrcodeKey: qrcode.Data.QrcodeKey,
	}, nil
}

// GetLoginInfo godoc
// @Summary     获取二维码状态
// @Description
// @Tags        bili
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       qrcode_key   	query    	string  	true  	"QrcodeKey"
// @Router      /bili/qrcode/getLoginInfo [get]
func GetLoginInfo(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	key := c.Query("qrcode_key")

	cookie, err := bilibot.GetLoginInfo(key)
	if err != nil {
		return nil, err
	}

	if err := bot.Add(cookie, id); err != nil {
		if err := bot.Update(cookie, id); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

// CheckLogin godoc
// @Summary     查询Bot登陆状态
// @Description
// @Tags        bili
// @Accept      json
// @Produce     json
// @Security 	ApiKeyAuth
// @Param       sessdata   	query    	string  	true  	"cookie当中的SESSDATA"
// @Success		200			{object}	api.BotInfo
// @Router      /bili/bot/check [get]
func CheckLogin(c *engine.Context) (interface{}, error) {
	sessdata := c.Query("sessdata")
	Bot, err := bilibot.GetBotInfo(&http.Cookie{Name: "SESSDATA", Value: sessdata})
	if err != nil {
		return nil, err
	}
	return api.BotInfo{
		BotID:   Bot.Data.Mid,
		Name:    Bot.Data.Name,
		IsLogin: Bot.Data.IsLogin,
		Face:    Bot.Data.Face,
	}, nil
}
