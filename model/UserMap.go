package model

type UserMap struct {
	Id uint `gorm:"primarykey"`

	FriendId uint

	UserId uint
}
