package bot

import (
	"net/http"
	"strconv"

	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/utils"
)

func getBotInfo(SESSDATA *http.Cookie) (*bilibot.BotInfo, error) {
	return bilibot.GetBotInfo(SESSDATA) // cookie[0] 表示 SESSDATA
}

func Add(cookie []*http.Cookie, UserID uint) error {
	botInfo, err := getBotInfo(cookie[0])
	if err != nil {
		return err
	}

	return dao.Create(&model.Bot{
		UID:     strconv.Itoa(int(botInfo.Data.Mid)),
		Name:    botInfo.Data.Name,
		Face:    botInfo.Data.Face,
		Cookie:  utils.CookieToString(cookie),
		IsLogin: botInfo.Data.IsLogin,
		UserID:  UserID,
	})
}

func Update(cookie []*http.Cookie, UserID uint) error {
	botInfo, err := getBotInfo(cookie[0])
	if err != nil {
		return err
	}

	return dao.Save(&model.Bot{
		UID:     strconv.Itoa(int(botInfo.Data.Mid)),
		Name:    botInfo.Data.Name,
		Face:    botInfo.Data.Face,
		Cookie:  utils.CookieToString(cookie),
		IsLogin: botInfo.Data.IsLogin,
		UserID:  UserID,
	})
}

func Get(uid string) (*model.Bot, error) {
	if uid == "" {
		return nil, e.ErrInvalidParam
	}
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

func Delete(uid string) error {
	return dao.Delete(&model.Bot{UID: uid})
}
