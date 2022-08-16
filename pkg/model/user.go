package model

import "github.com/Augenblick-tech/bilibot/lib/db"

type User struct {
	db.BaseModel
	Name     string `gorm:"type:varchar(255);not null;unique" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
}
