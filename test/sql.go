package main

import (
	"chatroom/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:8515367@tcp(192.168.0.6:3306)/chat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema

	db.AutoMigrate(&model.User{})

	db.AutoMigrate(&model.UserMessage{})

	db.AutoMigrate(&model.UserMap{})

	db.AutoMigrate(&model.Group{})
	db.AutoMigrate(&model.GroupMap{})

	db.AutoMigrate(&model.GroupMessage{})
}
