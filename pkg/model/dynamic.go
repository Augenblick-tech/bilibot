package model

import bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"

type Dynamic struct {
	DynamicID string `gorm:"type:text;not null;primary_key"`
	PubTS     uint64 `gorm:"type:integer;not null"`
	Content   string `gorm:"type:text;not null"`
	AuthorID  uint   `gorm:"type:integer;not null"`
}

func ToDynamic(dynamic ...bilibot.Dynamic) (dynamics []Dynamic) {
	for _, v := range dynamic {
		dynamics = append(dynamics, Dynamic{
			DynamicID: v.ID,
			PubTS:     uint64(v.Modules.Author.PubTS),
			Content:   v.Modules.Content.Desc.Text,
			AuthorID:  uint(v.Modules.Author.Mid),
		})
	}
	return
}
