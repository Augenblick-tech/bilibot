package model

type Author struct {
	UID   string `gorm:"type:integer;not null;primary_key" json:"author_id"`
	Name  string `gorm:"type:varchar(255);not null;unique" json:"name"`
	Face  string `gorm:"type:text" json:"face"`
	BotID string `gorm:"type:integer;not null" json:"bot_id"`
}
