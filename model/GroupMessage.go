package model

import "time"

type GroupMessage struct {
	Id uint `gorm:"primarykey"`

	UserId uint

	groupId uint

	content string

	messageType string

	SendTime time.Time
}
