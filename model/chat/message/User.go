package response

import "time"

type User struct {
	UserId     uint
	Username   string
	Avatar     string
	Role       string
	Tag        string
	Messages   []UserMessage
	CreateTime time.Time
}
