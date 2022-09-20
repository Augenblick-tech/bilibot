package main

import (
	"github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/conf"
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/lib/task"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/route"
)

// @title bilibot
// @version 2.0
// @description a bilibot server
// @termsOfService http://swagger.io/terms/

// @contact.name lonzzi
// @contact.url https://ronki.moe
// @contact.email lonzzi@qq.com

// @BasePath /v2

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
// @description					Bearer Token

func main() {
	conf.LoadDefaultConfig()
	InitDB()
	task.Start()
	if conf.C.Server.Domain == "" {
		docs.SwaggerInfo.Host = conf.C.Server.Addr
	} else {
		docs.SwaggerInfo.Host = conf.C.Server.Domain
	}
	route.Route(conf.C.Server.Addr)
}

func InitDB() {
	if err := db.Init(conf.C.DB.DbType, conf.C.DB.Data); err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(
		model.User{},
		model.Author{},
		model.Dynamic{},
		model.Bot{},
		model.Task{},
		model.Email{},
	); err != nil {
		panic(err)
	}
}
