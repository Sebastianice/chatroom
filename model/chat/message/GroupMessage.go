package response

import "time"

type GroupMessage struct {
	UserId      uint
	GroupId     uint
	Content     string
	Width       int
	Height      int
	MessageType string
	SendTime    time.Time
}
