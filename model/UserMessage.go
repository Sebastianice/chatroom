package model

import "time"

type UserMessage struct {
	Id uint `gorm:"primarykey"`

	UserId uint

	friendId uint

	content string

	messageType string

	SendTime time.Time
}
