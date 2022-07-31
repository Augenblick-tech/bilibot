package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lonzzi/bilibot/route"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	r := gin.Default()
	route.InitRoute(r)
	r.Run(":" + viper.GetString("server.port"))
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
