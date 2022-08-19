package main

import (
	"github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/conf"
	"github.com/Augenblick-tech/bilibot/lib/db"
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
func main() {
	conf.LoadDefaultConfig()
	InitDB()
	docs.SwaggerInfo.Host = conf.C.Server.Domain
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
	); err != nil {
		panic(err)
	}
}
