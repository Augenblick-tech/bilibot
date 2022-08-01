package route

import (
	"github.com/Augenblick-tech/bilibot/api"
	"github.com/Augenblick-tech/bilibot/api/bili"
	"github.com/gin-gonic/gin"
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
		v2.Handle("GET", "/dynamic/latest", api.GetLatestDynamic)
		v2.Handle("GET", "/dynamic/refresh", api.RefreshDynamic)
		v2.Handle("GET", "/dynamic/status", api.GetStatus)
		v2.Handle("GET", "/dynamic/stop", api.StopRefreshDynamic)
	}

	bi := v2.Group("/bili")
	{
		bi.Handle("GET", "/qrcode/getLoginUrl", bili.GetLoginUrl)
		bi.Handle("POST", "/login/getLoginInfo", bili.GetLoginInfo)
		bi.Handle("GET", "/dynamic/getDynamic", bili.GetDynamic)
		bi.Handle("POST", "/reply/add", bili.AddReply)
	}
}
