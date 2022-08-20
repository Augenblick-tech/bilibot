package bili

import (
	"net/http"

	"github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model/api"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
)

// GetLoginUrl godoc
// @Summary      获取二维码登录链接
// @Description
// @Tags         bili
// @Produce      json
// @Param 		 Authorization 	header 	string	true	"Bearer 用户令牌"
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
// @Tags         bili
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header	string				true	"Bearer 用户令牌"
// @Param        qrcode   		body    api.BiliQrCodeInfo  true  	"oauthKey"
// @Router       /bili/qrcode/getLoginInfo [post]
func GetLoginInfo(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	var oauth = api.BiliQrCodeInfo{}

	err := c.Bind(&oauth)
	if err != nil {
		return nil, err
	}

	cookie, err := bilibot.GetLoginInfo(oauth.OauthKey)
	if err != nil {
		return nil, err
	}

	return cookie, bot.Add(cookie, id)
}

// CheckLogin godoc
// @Summary      查询Bot登陆状态
// @Description
// @Tags         bili
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header	string				true	"Bearer 用户令牌"
// @Param        SESSDATA   	body    api.BiliAuthInfo  	true  	"SESSDATA"
// @Router       /bili/bot/check [post]
func CheckLogin(c *engine.Context) (interface{}, error) {
	var cookie = api.BiliAuthInfo{}

	err := c.Bind(&cookie)
	if err != nil {
		return nil, err
	}
	return bilibot.GetBotInfo(&http.Cookie{Name: "SESSDATA", Value: cookie.SESSDATA})
}
