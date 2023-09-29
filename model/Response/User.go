package response

import "time"

type UserMap struct {
	UserId     uint
	Username   string
	Avatar     string
	Role       string
	Tag        string
	Messages   []UserMapMessage
	CreateTime time.Time
}
