package model

import "chatroom/global"

type User struct {
	global.GVA_MODEL

	UserName string `json:"username" gorm:"username"`

	Password string `json:"password" gorm:"password"`

	Avatar string `json:"avatar" gorm:"avatar"`

	Role string `json:"role" gorm:"role"`

	Status string `json:"status" gorm:"status"`

	Tag string `json:"tag" gorm:"tag"`
}

func (u User) TableName() string {
	return "chat_user"
}
