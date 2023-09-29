package response

import "time"

type UserMapMessage struct {
	UserId      uint
	FriendId    uint
	Content     string
	Width       int
	Height      int
	MessageType string
	SendTime    time.Time
}
