package def

import (
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	Log = log.New()
)

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

type ResultSecKill struct {
	ProductId int
	UserId    string
	Mes       string
	Token     string
	Nance     string
}
