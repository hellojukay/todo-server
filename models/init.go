package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Store *gorm.DB
)

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
	Store = db
	log.Println("数据库初始化中 ...")
	AutoMigrate()
	log.Println("数据库初始化完成")
}

func AutoMigrate() {
	Store.AutoMigrate(&Task{})
}
