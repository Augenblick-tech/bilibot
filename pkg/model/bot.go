package model

type Bot struct {
	UID     string `gorm:"type:integer;not null;primary_key"`
	Name    string `gorm:"type:varchar(255);not null"`
	Face    string `gorm:"type:text"`
	Cookie  string `gorm:"type:text;not null"`
	IsLogin bool   `gorm:"type:boolean;not null"`
	UserID  uint   `gorm:"type:integer;not null"`
}
