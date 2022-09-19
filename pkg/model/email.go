package model

type Email struct {
	Host   string `gorm:"type:varchar(255)" json:"host"`
	Port   int    `gorm:"type:integer" json:"port"`
	From   string `gorm:"type:varchar(255)" json:"from"`
	To     string `gorm:"type:varchar(255)" json:"to"`
	Pass   string `gorm:"type:varchar(255)" json:"pass"`
	UserID uint   `gorm:"type:integer;not null" json:"user_id"`
}
