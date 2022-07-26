package route

import (
	"github.com/gin-gonic/gin"
	"github.com/lonzzi/BiliUpDynamicBot/api"
	"github.com/lonzzi/BiliUpDynamicBot/api/bili"
)

func InitRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v2 := r.Group("/v2")
	{
		v2.Handle("POST", "/login", api.Login)
	}

	bilibili := v2.Group("/bili")
	{
		bilibili.Handle("GET", "/qrcode/getLoginUrl", bili.GetLoginUrl)
		bilibili.Handle("POST", "/login/getLoginInfo", bili.GetLoginInfo)
		bilibili.Handle("GET", "/dynamic/getDynamic", bili.GetDynamic)
	}
}
