package bili

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/bilibot/pkg/model"
	"github.com/lonzzi/bilibot/response"
	"github.com/spf13/viper"
)

func GetLoginUrl(c *gin.Context) {
	var r response.Response

	qrcode, err := model.GetLoginUrl()
	if err != nil {
		r.Code = response.CodeBiliLoginError
		r.JSON(c, http.StatusBadGateway, err.Error(), nil)
		return
	}

	r.JSON(c, http.StatusOK, "success", qrcode)
}

func GetLoginInfo(c *gin.Context) {
	var r response.Response

	account, err := model.GetLoginInfo(c.PostForm("oauthKey"), 60)
	if err != nil {
		r.Code = response.CodeBiliLoginError
		r.JSON(c, http.StatusRequestTimeout, err.Error(), nil)
		return
	}

	viper.Set("account.SESSDATA", account.SESSDATA)
	viper.Set("account.bili_jct", account.BiliJct)

	r.JSON(c, http.StatusOK, "success", account)
}
