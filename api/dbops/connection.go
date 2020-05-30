package dbops

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	pool redis.Pool
	db   *gorm.DB
	err  error
)

func init() {
	//mysql db init
	db, err = gorm.Open("mysql", "pace:123@tcp(127.0.0.1:3306)/piliVideo?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	//redis pool init
	pool = redis.Pool{
		MaxIdle:     16,
		MaxActive:   32,
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}
