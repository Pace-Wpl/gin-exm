package def

import (
	log "github.com/sirupsen/logrus"
)

const (
	//默认头像dir
	DEFAULT_ICON = "../icon/default.jpg"
)

var (
	Log = log.New()
)

type RespMes struct {
	Code int    `json:"code"`
	Mes  string `json:"message"`
}

type ReqModifyPwd struct {
	Pwd    string `json:"password"`
	NewPwd string `json:"new_password"`
}

type ReqUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username string `json:"name"`
	Pwd      string `json:"password"`
	Icon     string `json:"icon"`
}

type Session struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
