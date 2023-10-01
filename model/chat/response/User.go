package response

import "chatroom/model"

type UserResponse struct {
	User  model.User `json:"user"`
	Token string     `json:"token"`
}
