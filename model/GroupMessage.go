package model

import "time"

type GroupMessage struct {
	Id uint `gorm:"primarykey"`

	UserId uint `gorm:"userId" json:"userId"`

	GroupId uint `gorm:"groupId" json:"groupId"`

	Content string `gorm:"content" json:"content"`

	MessageType string `gorm:"messageType" json:"messageType"`

	SendTime time.Time `gorm:"sendTime" json:"time"`
}

func (u GroupMessage) TableName() string {
	return "chat_group_message"
}
