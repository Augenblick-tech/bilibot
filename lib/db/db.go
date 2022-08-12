package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

type DbType int

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

const (
	DB_Nil DbType = iota
	SQLite3
	Mysql
)

var DB *Database

func Init(t DbType, d string) error {
	DB = &Database{}
	switch t {
	case SQLite3:
		{
			db, err := gorm.Open(sqlite.Open(d), &gorm.Config{})
			if err != nil {
				return err
			}
			DB.DB = db
		}
	case Mysql:
		{

		}
	default:
		{
			return fmt.Errorf("cannot init db, unknow db type %v", t)
		}
	}
	return nil
}

func AutoMigrate(dst ...interface{}) error {
	return DB.AutoMigrate(dst)
}

func Transaction(f func(tdb *Database) error) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		db := &Database{
			DB: tx,
		}
		return f(db)
	})
}
