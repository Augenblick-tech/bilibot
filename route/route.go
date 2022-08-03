package route

import (
	"github.com/Augenblick-tech/bilibot/api"
	"github.com/Augenblick-tech/bilibot/api/bili"
	_ "github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Route(addr string) {
	engine.SetMode("debug")
	e := engine.NewDefaultEngine()
	e.Use(engine.Result)

	e.GET("/ping", "pong", func(ctx *engine.Context) (interface{}, error) {
		return "pong", nil
	})

	e.GET("/swagger/*any", "swagger", func(ctx *engine.Context) (interface{}, error) {
		gs.WrapHandler(swaggerFiles.Handler)(ctx.Context)
		return nil, nil
	})

	v2 := e.Group("/v2")
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
