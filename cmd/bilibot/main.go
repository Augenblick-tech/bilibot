package main

import (
	"log"

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

// @host localhost:2333
// @BasePath /v2
func main() {
	InitConfig()
	route.Route(viper.GetString("server.addr"))
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
