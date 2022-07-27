package bilibot

import (
	"io"
	"log"
	"os"
	"time"

	"github.com/lonzzi/BiliUpDynamicBot/pkg/model"
	"github.com/lonzzi/BiliUpDynamicBot/pkg/utils"
	"github.com/spf13/viper"
)

func BotLogin() {

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

	qrCodeResponse, err := model.GetLoginUrl()
	if err != nil {
		log.Fatal(err)
	}
	account, err := model.GetLoginInfo(qrCodeResponse.Data.OauthKey, 60)
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

		dynamic, err := model.GetDynamic(Mid)
		if err != nil {
			log.Fatal(err)
		}

		// 获取历史数据
		oldDynamics, err := model.GetHistoryDynamics(viper.GetString("HistroyDynamicPath"))
		if err != nil {
			log.Fatal(err)
		}

		// 判断是否有新的动态
		if model.IsDynamicExist(oldDynamics, dynamic[0]) {
			if !isPrinted {
				log.Println("无新的动态")
				log.Println("上一条动态：", dynamic[0].Modules.Content.Desc.Text)
				isPrinted = true
			}
			log.Println("等待动态更新中...")
			continue
		}

		log.Println("有新的动态：", dynamic[0].Modules.Content.Desc.Text)
		_, err = model.AddNewDynamic(oldDynamics, dynamic[0])
		if err != nil {
			log.Println("添加动态错误：", err)
			continue
		}

		threshold := 0.1
		message, err := utils.UnicodeToStr(dynamic[0].Modules.Content.Desc.Text, threshold)
		if err != nil {
			log.Println("转换为字符错误：", err)
			message = utils.StrToUnicode(dynamic[0].Modules.Content.Desc.Text)
			log.Println("转换为Unicode：", message)
		}
		message = viper.GetString("uploader.MessageHead") + "\n" + message + "\n" + viper.GetString("uploader.MessageTail")
		ReplyResponse, err := model.DynamicReply(dynamic[0], message)
		if err != nil {
			log.Println("回复出错：", err)
		}

		if ReplyResponse != nil {
			log.Println(ReplyResponse.Data.SuccessToast)
			log.Println("发送内容：\n", message)
		} else {
			log.Println("ReplyResponse is nil")
		}
	}
}