package dao

import "github.com/Augenblick-tech/bilibot/lib/db"

func Create(obj interface{}) error {
	return db.DB.Create(obj).Error
}

func First(obj interface{}) error {
	return db.DB.First(&obj).Error
}

func Delete(obj interface{}) error {
	return db.DB.Delete(obj).Error
}
