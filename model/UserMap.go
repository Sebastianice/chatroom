package model

type UserMap struct {
	Id uint `gorm:"primarykey"`

	FriendId uint `gorm:"friendId" json:"friendId"`

	UserId uint `gorm:"userId" json:"userId"`
}

func (u UserMap) TableName() string {
	return "chat_user_map"
}
