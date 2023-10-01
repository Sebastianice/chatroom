package chat

import (
	"chatroom/global"
	"chatroom/model"
	"chatroom/model/chat/request"
	"chatroom/utils"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
}

func (u *UserService) Register(user request.RegisterUserParam) (err error) {

	if errors.Is(global.GVA_DB.Where("name= ?", user.UserName).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同用户名")
	}

	return global.GVA_DB.Create(&model.User{
		UserName: user.UserName,
		Password: utils.Md5(user.Password),
	}).Error
}

func (u *UserService) Login(params request.UserLoginParam) (err error, user model.User) {
	err = global.GVA_DB.Where("name=? AND password= ?", params.UserName, params.PasswordMd5).First(&user).Error
	if user != (model.User{}) {

	}
	return err, user
}
