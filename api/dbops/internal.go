package dbops

import "github.com/gomodule/redigo/redis"

// 0代表不在map中，1代表白名单，2代表黑名单
func IpMap(ip string) int {
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.Int(conn.Do("hget", "ip_map", ip))
	if err != nil {
		return 0
	} else {
		return res
	}
}

// 0代表不在map中，1代表白名单，2代表黑名单
func IdMap(id string) int {
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.Int(conn.Do("hget", "id_map", id))
	if err != nil {
		return 0
	} else {
		return res
	}
}

// refer白名单判断
func ReferWhiteMap(refer string) bool {
	conn := pool.Get()
	defer conn.Close()
	res, err := redis.Bool(conn.Do("hexists", "refer_white_map", refer))
	if err != nil {
		return false
	} else {
		return res
	}
}
