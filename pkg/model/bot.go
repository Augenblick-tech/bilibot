package model

type Bot struct {
	UID     string `gorm:"type:integer;not null;primary_key" json:"bot_id"`
	Name    string `gorm:"type:varchar(255);not null" json:"name"`
	Face    string `gorm:"type:text" json:"face"`
	Cookie  string `gorm:"type:text;not null" json:"cookie"`
	IsLogin bool   `gorm:"type:boolean;not null" json:"is_login"`
	UserID  uint   `gorm:"type:integer;not null" json:"user_id"`
}
