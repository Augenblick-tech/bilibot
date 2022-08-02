package bili

import (
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/spf13/viper"
)

func GetLoginUrl(c *engine.Context) (interface{}, error) {
	qrcode, err := model.GetLoginUrl()
	if err != nil {
		return nil, err
	}

	return qrcode, nil
}

func GetLoginInfo(c *engine.Context) (interface{}, error) {
	account, err := model.GetLoginInfo(c.Query("oauthKey"), 60)
	if err != nil {
		return nil, err
	}

	viper.Set("account.SESSDATA", account.SESSDATA)
	viper.Set("account.bili_jct", account.BiliJct)

	return account, nil
}
