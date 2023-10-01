package model

import "time"

type UserMessage struct {
	Id uint `gorm:"primarykey"`

	UserId uint `gorm:"userId" json:"userId" `

	FriendId uint `gorm:"friendId" json:"friendId"`

	Content string `gorm:"content" json:"content"`

	//MessageType string `gorm:"messagetype" json:"friendId"`

	SendTime time.Time `gorm:"sendTime" json:"time"`
}

func (u UserMessage) TableName() string {
	return "chat_user_message"
}
