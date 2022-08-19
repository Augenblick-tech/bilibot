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

type oauthInfo struct {
	OauthKey string `json:"oauthKey" binding:"required"`
}
// GetLoginInfo godoc
// @Summary      获取二维码状态
// @Description
// @Tags         bili
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header	string		true	"Bearer 用户令牌"
// @Param        qrcode   		body    oauthInfo  	true  	"oauthKey"
// @Router       /bili/qrcode/getLoginInfo [post]
func GetLoginInfo(c *engine.Context) (interface{}, error) {
	id := c.Context.GetUint("UserID")
	var oauth = oauthInfo{}

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

type cookieInfo struct {
	SESSDATA string `json:"SESSDATA" binding:"required"`
}
// CheckLogin godoc
// @Summary      查询Bot登陆状态
// @Description
// @Tags         bili
// @Accept       json
// @Produce      json
// @Param 		 Authorization 	header	string		true	"Bearer 用户令牌"
// @Param        SESSDATA   	body    cookieInfo  true  	"SESSDATA"
// @Router       /bili/bot/check [post]
func CheckLogin(c *engine.Context) (interface{}, error) {
	var cookie = cookieInfo{}

	err := c.Bind(&cookie)
	if err != nil {
		return nil, err
	}
	return bilibot.GetBotInfo(&http.Cookie{Name: "SESSDATA", Value: cookie.SESSDATA})
}
