package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"username" binding:"required"`
	Password string `gorm:"type:varchar(255);not null" json:"password" binding:"required"`
}

func (u *User) Create() error {
	return DB.Create(u).Error
}

func (u *User) Find(names ...string) error {
	if len(names) <= 0 {
		return DB.First(u).Error
	}

	return DB.Where("name in (?)", names).First(u).Error
}
