package model

import (
	"errors"

	"github.com/Augenblick-tech/bilibot/pkg/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(255);not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

var DB *gorm.DB

func init() {
	DB = db.Get()
}

func (u *User) Create() error {
	return DB.Create(u).Error
}

func (u *User) Get(names ...string) error {
	if len(names) <= 0 {
		return DB.First(u).Error
	}

	result := DB.Where("name in (?)", names).First(u)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}
