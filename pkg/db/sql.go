package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("./db/bilibot.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{}, &Bot{}, &Author{}, &Dynamic{})

	DB = db
}

func Get() *gorm.DB {
	return DB
}
