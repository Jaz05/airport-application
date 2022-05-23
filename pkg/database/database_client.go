package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetClient() *gorm.DB {
	if db == nil {
		db = connectToDB()
	}
	return db
}

func connectToDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/airport?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
