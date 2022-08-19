package model

type Author struct {
	UID   string `gorm:"type:integer;not null;primary_key"`
	Name  string `gorm:"type:varchar(255);not null;unique"`
	Face  string `gorm:"type:text"`
	BotID string `gorm:"type:integer;not null"`
}
