package model

type Author struct {
	UID   string `gorm:"type:integer;not null;primary_key"`
	Name  string `gorm:"type:varchar(255);not null"`
	Face  string `gorm:"type:text"`
	BotID uint   `gorm:"type:integer;not null"`
	Bot   Bot    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
