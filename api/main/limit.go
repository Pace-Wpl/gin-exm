package main

import (
	"sync"

	"github.com/gin-exm/api/def"
)

var (
	Limit = &LimitStr{
		UserLimitMap: make(map[string]*SecLimit),
		IpLimitMap:   make(map[string]*MinLimit),
	}
)

type LimitStr struct {
	UserLimitMap map[string]*SecLimit
	IpLimitMap   map[string]*MinLimit
	lock         sync.Mutex
}

type SecLimit struct {
	count   int
	curTime int64
}

type MinLimit struct {
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

func (p *MinLimit) calculate(nowTime int64) int {
	if nowTime-p.curTime >= 60 {
		p.count = 1
		p.curTime = nowTime
		return p.count
	}

	p.count++
	return p.count
}

//访问限制
func Antispam(req *def.ReqSecKill) bool {
	if userAndIpFilter(req) || referFilter(req) {
		return true
	}
	return false
}

func userAndIpFilter(req *def.ReqSecKill) bool {
	/**   黑、白名单控制  **/
	Limit.lock.Lock()
	//uid 频率控制
	limit, ok := Limit.UserLimitMap[req.UserID]
	if !ok {
		limit = &SecLimit{}
		Limit.UserLimitMap[req.UserID] = limit
	}
	secIDCount := limit.calculate(req.AccessTime.Unix())

	//ip 频率控制
	limitM, ok := Limit.IpLimitMap[req.ClientIp]
	if !ok {
		limitM = &MinLimit{}
		Limit.IpLimitMap[req.ClientIp] = limitM
	}
	secIPCount := limitM.calculate(req.AccessTime.Unix())

	Limit.lock.Unlock()
	if secIDCount > def.Conf.UserSecAccessLimit {
		def.Log.Warnln("用户id:" + req.UserID + "访问过多")
		return true
	} else if secIPCount > def.Conf.IpSecAccessLimit {
		def.Log.Warnln("ip:" + req.ClientIp + "访问过多")
		return true
	}
	return false
}

//refer过滤
func referFilter(req *def.ReqSecKill) bool {
	/**   白名单控制  **/
	return false
}
