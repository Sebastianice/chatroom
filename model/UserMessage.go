package model

import "time"

type UserMapMessage struct {
	Id uint `gorm:"primarykey"`

	UserId uint

	friendId uint

	content string

	messageType string

	SendTime time.Time
}
