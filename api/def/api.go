package def

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	PRODUCT_STATUS_NOMAL   = 0
	PRODUCT_STATUS_SELLOUT = 1
	PRODUCT_STATUS_END     = 2
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

type ReqSecKill struct {
	ProductID  int
	UserID     string
	Sourct     string
	AuthCode   string
	Time       string
	Nance      string
	AccessTime time.Time
}

type User struct {
	Username string `json:"name"`
	Pwd      string `json:"password"`
	Icon     string `json:"icon"`
}

type Session struct {
	UserId string `json:"user_id"`
	ID     string `json:"id"`
}
