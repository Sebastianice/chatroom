package model

import "chatroom/global"

type Group struct {
	global.GVA_MODEL

	UserId    uint   `gorm:"userId" json:"userId"`
	GroupName string `gorm:"groupName" json:"groupName"`
	Notice    string `gorm:"notice" json:"notice"`
}

func (u Group) TableName() string {
	return "chat_group"
}
