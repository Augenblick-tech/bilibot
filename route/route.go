package route

import (
	"github.com/Augenblick-tech/bilibot/api/bili"
	"github.com/Augenblick-tech/bilibot/api/dynamic"
	"github.com/Augenblick-tech/bilibot/api/web"
	_ "github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/conf"
	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/lib/jwt"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func Route(addr string) {
	engine.SetMode(conf.C.Server.Mode)
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
		v2.POST("/web/register", "register", web.Register)
		v2.POST("/web/login", "login", web.Login)
		v2.GET("/web/refreshToken", "refreshToken", web.RefreshToken)
	}

	webs := v2.Group("/web").Use(jwt.JWTAuth)
	{
		webs.GET("/bot/list", "getBotList", web.GetBotList)
		webs.GET("/author/list", "getAuthorList", web.GetAuthorList)
		webs.POST("/author/add", "addAuthor", web.AddAuthor)
		webs.POST("/setting/update", "UpdateSetting", web.UpdateSettings)
		webs.POST("/task/update", "UpdateTaskStatus", web.SetTaskStatus)
		webs.GET("/task/status", "GetTaskStatus", web.GetTaskStatus)
		dynm := webs.Group("/dynamic")
		{
			dynm.GET("/list", "getDynamicList", web.GetDynamicList)
			dynm.GET("/latest", "getLatestDynamic", dynamic.Latest)
			dynm.GET("/listen", "listenDynamic", dynamic.Listen)
			dynm.GET("/status", "getStatus", dynamic.Status)
			dynm.GET("/stop", "stopRefreshDynamic", dynamic.Stop)
		}
	}

	bi := v2.Group("/bili").Use(jwt.JWTAuth)
	{
		bi.GET("/qrcode/getLoginUrl", "getLoginUrl", bili.GetLoginUrl)
		bi.GET("/qrcode/getLoginInfo", "getLoginInfo", bili.GetLoginInfo)
		bi.GET("/bot/check", "checkLogin", bili.CheckLogin)
		bi.GET("/dynamic/getDynamic", "getDynamic", bili.GetDynamic)
		bi.POST("/reply/add", "addReply", bili.AddReply)
	}

	e.Run(addr)
}
