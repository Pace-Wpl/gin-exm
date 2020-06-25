package def

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	PRODUCT_STATUS_NOMAL   = 0
	PRODUCT_STATUS_SELLOUT = 1
	PRODUCT_STATUS_END     = 2
	PRODUCT_ACTIVITY_PRE   = 0
	PRODUCT_ACTIVITY_BEGIN = 1
	PRODUCT_ACTIVITY_END   = 2
)

var (
	Log = log.New()
)

type RespMes struct {
	Code int    `json:"code"`
	Mes  string `json:"message"`
}

type RespProductInfo struct {
	ProductID int   `json:"product_id"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
	Status    int   `json:"status"`
	Activity  int   `json:"activity"`
	Total     int   `json:"total"`
}

type RespSecKillProduct struct {
	UserID    string `json:"user_id"`
	ProductID int    `json:"product_id"`
	Mes       string `json:"message"`
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
	ProductID   int
	UserID      string
	Source      string
	AuthCode    string
	Time        string
	Nance       string
	ClientIp    string
	CLientRefer string
	AccessTime  time.Time
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

type ResultSecKill struct {
	ProductId int
	UserId    string
	Mes       string
	Token     string
}
