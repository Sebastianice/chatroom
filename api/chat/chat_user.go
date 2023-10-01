package chat

import (
	"chatroom/global"
	"chatroom/model"
	"chatroom/model/chat/request"
	"chatroom/model/chat/response"
	common "chatroom/model/commom/response"
	"chatroom/service"
	"chatroom/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserApi struct {
}

func (u *UserApi) GetUserID() {

}

func (u *UserApi) Login(c *gin.Context) {
	var user request.UserLoginParam
	err := c.ShouldBindJSON(&user)
	if err != nil {
		return
	}
	err, m := service.ServiceGroupApp.ChatServiceGroup.UserService.Login(user)
	u.TokenNext(c, m)
}

func (u *UserApi) Register(c *gin.Context) {

}

// TokenNext 登录以后签发jwt
func (u *UserApi) TokenNext(c *gin.Context, user model.User) {

	claims := utils.CreateClaims(request.BaseClaims{
		UserName: user.UserName,
		Password: user.Password,
	})
	token, err := utils.CreateToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		common.FailWithMessage("获取token失败", c)
		return
	}

	common.OkWithDetailed(response.UserResponse{
		User:  user,
		Token: token,
	}, "登录成功", c)
}
