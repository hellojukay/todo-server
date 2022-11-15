package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	Store *gorm.DB
)

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
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
