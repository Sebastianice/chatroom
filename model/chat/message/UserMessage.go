package message

import "time"

type UserMessage struct {
	UserId      uint
	FriendId    uint
	Content     string
	Width       int
	Height      int
	MessageType string
	SendTime    time.Time
}
