package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("./pkg/db/bilibot.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db
}

func Get() *gorm.DB {
	return DB
}