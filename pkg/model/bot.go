package model

import bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"

type Bot struct {
	UID     uint   `gorm:"type:integer;not null;primary_key"`
	Name    string `gorm:"type:varchar(255);not null"`
	Face    string `gorm:"type:text"`
	Cookie  string `gorm:"type:text;not null"`
	IsLogin bool   `gorm:"type:boolean;not null"`
	UserID  uint   `gorm:"type:integer;not null"`
	User    User   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func ToBot(bot bilibot.BotInfo, cookie string, UserID uint) Bot {
	return Bot{
		UID:     bot.Data.Mid,
		Name:    bot.Data.Name,
		Face:    bot.Data.Face,
		Cookie:  cookie,
		IsLogin: bot.Data.IsLogin,
		UserID:  UserID,
	}
}
