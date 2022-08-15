package bot

import (
	"net/http"

	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/utils"
)

func Add(cookie []*http.Cookie, UserID uint) error {
	botInfo, err := bilibot.GetBotInfo(cookie[3]) // cookie[3] 表示 SESSDATA
	if err != nil {
		return err
	}

	return dao.Create(&model.Bot{
		UID:     botInfo.Data.Mid,
		Name:    botInfo.Data.Name,
		Face:    botInfo.Data.Face,
		Cookie:  utils.CookieToString(cookie),
		IsLogin: botInfo.Data.IsLogin,
		UserID:  UserID,
	})
}

func Get(uid uint) (*model.Bot, error) {
	bot := model.Bot{
		UID: uid,
	}
	if err := dao.First(&bot); err != nil {
		return nil, err
	}
	return &bot, nil
}

func GetList(UserID uint) ([]*model.Bot, error) {
	return dao.GetBotList(UserID)
}

func Delete(uid uint) error {
	return dao.Delete(&model.Bot{UID: uid})
}
