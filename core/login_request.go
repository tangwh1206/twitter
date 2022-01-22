package core

type LoginReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type RegisterReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
