package dbops

import (
	"time"

	"github.com/gin-exm/api/def"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	pool redis.Pool
	db   *gorm.DB
	err  error
)

func InitDB() error {
	//mysql db init
	str := def.Conf.Mysql.User + ":" + def.Conf.Mysql.Pwd + "@tcp(" + def.Conf.Mysql.Addr + ")/" +
		def.Conf.Mysql.Database + "?" + def.Conf.Mysql.Config
	def.Log.Infoln(str)
	db, err = gorm.Open("mysql", str)
	if err != nil {
		panic(err.Error())
	}

	//redis pool init
	pool = redis.Pool{
		MaxIdle:     def.Conf.Redis.MaxIdle,
		MaxActive:   def.Conf.Redis.MaxActive,
		IdleTimeout: time.Duration(def.Conf.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", def.Conf.Redis.Addr)
		},
	}
	_, err = pool.Get().Do("ping")
	if err != nil {
		panic(err.Error())
	}

	return nil
}
