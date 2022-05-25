package database

import (
	"airport/pkg/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var inMemoryDB *gorm.DB

func GetClient() *gorm.DB {
	if db == nil {
		db = connectToDB()
	}
	return db
}

func GetInMemoryClient() *gorm.DB {
	if inMemoryDB == nil {
		inMemoryDB = connectToInMemoryClient()
	}
	return inMemoryDB
}

func connectToDB() *gorm.DB {
	var conf = config.GetDBConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func connectToInMemoryClient() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
