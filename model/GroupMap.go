package model

import "chatroom/global"

type GroupMap struct {
	global.GVA_MODEL

	GroupId uint `gorm:"groupId" json:"groupId"`
	UserId  uint `gorm:"userId" json:"userId"`
}

func (u GroupMap) TableName() string {
	return "chat_group_map"
}
