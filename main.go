package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/lonzzi/BiliUpDynamicBot/e"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	_, err := os.Stat("./logs")
	if err != nil {
		os.Mkdir("./logs", os.ModePerm)
	}
	logFile, err := os.OpenFile("./logs/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	ticker := time.NewTicker(time.Duration(viper.GetInt("user.RefreshTime")) * time.Second)
	cnt := 1
	for {
		Mid := viper.GetString("uploader.mid")
		if Mid == "" {
			log.Println("Mid is empty")
			return
		}
		dynamic := GetLatestDynamic(Mid)

		<-ticker.C
		log.Println("运行次数：", cnt)
		cnt++
		message, err := Unicode2Str(dynamic.Content.Desc.Text, 0.1)
		if err != nil {
			log.Println(err)
			continue
		}
		err = MakeReply(*dynamic, message)
		if err != nil {
			log.Println(e.MakeReplyError, ":", err)
		}
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
