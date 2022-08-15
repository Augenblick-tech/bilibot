package dao

import (
	"github.com/Augenblick-tech/bilibot/lib/db"
	"gorm.io/gorm/clause"
)

func Create(obj interface{}) error {
	return db.DB.Create(obj).Error
}

func CreateWithIgonreConflict(obj interface{}) error {
	return db.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(obj).Error
}

func First(obj interface{}) error {
	return db.DB.First(&obj).Error
}

func Delete(obj interface{}) error {
	return db.DB.Delete(obj).Error
}

func Save(obj interface{}) error {
	return db.DB.Save(obj).Error
}
