package main

import (
	"log"

	"github.com/Augenblick-tech/bilibot/route"
	"github.com/spf13/viper"
)

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
