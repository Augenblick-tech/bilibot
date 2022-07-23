package main

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
)

func main() {
	InitConfig()
	isPrinted := false // 选择是否打印上一条动态
	viper.Set("HistroyDynamicPath", "./dynamic_history.json")
	Mid := viper.GetString("uploader.mid")
	if Mid == "" {
		log.Fatal("Mid is empty")
	}

	_, err := os.Stat("./logs")
	if err != nil {
		os.Mkdir("./logs", os.ModePerm)
	}

	account, err := Login()
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("account.SESSDATA", account.SESSDATA)
	viper.Set("account.bili_jct", account.BiliJct)

	logFile, err := os.OpenFile("./logs/"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	log.Printf("%d 秒后开始获取动态！", viper.GetInt("user.RefreshTime"))
	ticker := time.NewTicker(time.Duration(viper.GetInt("user.RefreshTime")) * time.Second)
	for {
		<-ticker.C

		dynamic, err := GetLatestDynamic(Mid)
		if err != nil {
			log.Fatal(err)
		}

		// 获取历史数据
		oldDynamics, err := GetHistoryDynamics(viper.GetString("HistroyDynamicPath"))
		if err != nil {
			log.Fatal(err)
		}

		// 判断是否有新的动态
		if IsExistDynamic(oldDynamics, *dynamic) {
			if !isPrinted {
				log.Println("无新的动态")
				log.Println("上一条动态：", dynamic.Content.Desc.Text)
				isPrinted = true
			}
			log.Println("等待动态更新中...")
			continue
		}

		log.Println("有新的动态：", dynamic.Content.Desc.Text)
		oldDynamics, err = AddNewDynamic(oldDynamics, *dynamic)
		if err != nil {
			log.Println("添加动态错误：", err)
			continue
		}

		threshold := 0.1
		message, err := Unicode2Str(dynamic.Content.Desc.Text, threshold)
		if err != nil {
			log.Println("动态转换错误: ", err, "阈值为：", threshold)
			isPrinted = false
			continue
		}
		commentResponse, err := MakeReply(oldDynamics, *dynamic, message)
		if err != nil {
			log.Println("回复出错：", err)
		}

		if commentResponse != nil {
			log.Println(commentResponse.Data.SuccessToast)
			log.Println("发送内容：", message)
		} else {
			log.Println("commentResponse is nil")
		}
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
