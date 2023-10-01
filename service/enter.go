package service

import "chatroom/service/chat"

type ServiceGroup struct {
	ChatServiceGroup chat.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
