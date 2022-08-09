package route

import (
	"github.com/Augenblick-tech/bilibot/api"
	"github.com/Augenblick-tech/bilibot/api/bili"
	"github.com/Augenblick-tech/bilibot/api/dynamic"
	_ "github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Route(addr string) {
	engine.SetMode("debug")
	e := engine.NewDefaultEngine()

	e.GET("/ping", "pong", func(ctx *engine.Context) (interface{}, error) {
		return "pong", nil
	})

	e.GET("/swagger/*any", "swagger", func(ctx *engine.Context) (interface{}, error) {
		gs.WrapHandler(swaggerFiles.Handler)(ctx.Context)
		return nil, nil
	})

	v2 := e.Group("/v2").Use(engine.Result)
	{
		v2.POST("/register", "register", api.Register)
		v2.POST("/login", "login", api.Login)
		v2.GET("/dynamic/latest", "getLatestDynamic", dynamic.Latest)
		v2.GET("/dynamic/listen", "listenDynamic", dynamic.Listen)
		v2.GET("/dynamic/status", "getStatus", dynamic.Status)
		v2.GET("/dynamic/stop", "stopRefreshDynamic", dynamic.Stop)
	}

	bi := v2.Group("/bili")
	{
		bi.GET("/qrcode/getLoginUrl", "getLoginUrl", bili.GetLoginUrl)
		bi.POST("/qrcode/getLoginInfo", "getLoginInfo", bili.GetLoginInfo)
		bi.GET("/dynamic/getDynamic", "getDynamic", bili.GetDynamic)
		bi.POST("/reply/add", "addReply", bili.AddReply)
	}

	e.Run(addr)
}
