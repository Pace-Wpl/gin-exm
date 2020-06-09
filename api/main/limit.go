package main

import (
	"sync"

	"github.com/gin-exm/api/def"
)

var (
	Limit = &LimitStr{
		UserLimitMap: make(map[string]*SecLimit),
		IpLimitMap:   make(map[string]*SecLimit),
	}
)

type LimitStr struct {
	UserLimitMap map[string]*SecLimit
	IpLimitMap   map[string]*SecLimit
	lock         sync.Mutex
}

type SecLimit struct {
	count   int
	curTime int64
}

func (p *SecLimit) calculate(nowTime int64) int {
	if p.curTime != nowTime {
		p.count = 1
		p.curTime = nowTime
		return p.count
	}

	p.count++
	return p.count
}

//访问限制
func AntiSpam(req *def.ReqSecKill) bool {
	Limit.lock.Lock()
	//uid 频率控制
	limit, ok := Limit.UserLimitMap[req.UserID]
	if !ok {
		limit = &SecLimit{}
		Limit.UserLimitMap[req.UserID] = limit
	}

	secIDCount := limit.calculate(req.AccessTime.Unix())

	if secIDCount > def.Conf.UserSecAccessLimit {
		return false
	}

	return true
}
