package model

import "chatroom/global"

type User struct {
	global.GVA_MODEL

	UserName string `json:"username"`

	Password string `json:"password"`

	Avatar string

	Role string

	Status string

	Tag string
}
