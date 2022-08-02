package route

import (
	"github.com/Augenblick-tech/bilibot/api"
	"github.com/Augenblick-tech/bilibot/api/bili"
	_ "github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
)

func Route(addr string) {
	engine.SetMode("debug")
	e := engine.NewDefaultEngine()

	e.GET("/ping", "pong", func(ctx *engine.Context) (interface{}, error) {
		return gin.H{
			"message": "pong",
		}, nil
	})

	e.GET("/swagger/*any", "swagger", func(ctx *engine.Context) (interface{}, error) {
		gs.WrapHandler(swaggerFiles.Handler)(ctx.Context)
		return "success", nil
	})

	v2 := e.Group("/v2").Use(engine.Result)
	{
		v2.POST("/login", "login", api.Login)
		v2.GET("/dynamic/latest", "getLatestDynamic", api.GetLatestDynamic)
		v2.GET("/dynamic/refresh", "refreshDynamic", api.RefreshDynamic)
		v2.GET("/dynamic/status", "getStatus", api.GetStatus)
		v2.GET("/dynamic/stop", "stopRefreshDynamic", api.StopRefreshDynamic)
	}

	bi := v2.Group("/bili")
	{
		bi.GET("/qrcode/getLoginUrl", "getLoginUrl", bili.GetLoginUrl)
		bi.GET("/login/getLoginInfo", "getLoginInfo", bili.GetLoginInfo)
		bi.GET("/dynamic/getDynamic", "getDynamic", bili.GetDynamic)
		bi.POST("/reply/add", "addReply", bili.AddReply)
	}

	e.Run(addr)
}
