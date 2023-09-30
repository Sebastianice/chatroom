package model

import "chatroom/global"

type GroupName struct {
	global.GVA_MODEL

	GroupId string
	UserId  string
}
