package request

type RegisterUserParam struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type UserLoginParam struct {
	UserName    string `json:"username"`
	PasswordMd5 string `json:"password"`
}
