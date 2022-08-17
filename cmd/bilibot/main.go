package main

import (
	"log"

	"github.com/Augenblick-tech/bilibot/docs"
	"github.com/Augenblick-tech/bilibot/lib/db"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/route"
	"github.com/spf13/viper"
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
	InitConfig()
	InitDB()
	docs.SwaggerInfo.Host = viper.GetString("server.domain") + ":" + viper.GetString("server.port")
	route.Route(viper.GetString("server.addr") + ":" + viper.GetString("server.port"))
}

func InitDB() {
	if err := db.Init(db.DbType(viper.GetInt("db.type")), viper.GetString("db.data")); err != nil {
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

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./conf")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
