package message

import "chatroom/global"

type Group struct {
	global.GVA_MODEL
	UserId    uint // 群主id
	GroupName string
	Notice    string
	Messages  []GroupMessage
}
