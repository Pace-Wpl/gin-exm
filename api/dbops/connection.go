package dbops

import (
	"time"

	etcd "github.com/etcd-io/etcd/clientv3"
	"github.com/gin-exm/api/def"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	etcdClient *etcd.Client
	pool       *redis.Pool
	db         *gorm.DB
	err        error
)

//mysql db init
func InitDB() error {
	str := def.Conf.Mysql.User + ":" + def.Conf.Mysql.Pwd + "@tcp(" + def.Conf.Mysql.Addr + ")/" +
		def.Conf.Mysql.Database + "?" + def.Conf.Mysql.Config
	def.Log.Debugln(str)
	db, err = gorm.Open("mysql", str)
	if err != nil {
		return err
	}

	return nil
}

//redis pool init
func InitRedis() error {
	pool = &redis.Pool{
		MaxIdle:     def.Conf.Redis.MaxIdle,
		MaxActive:   def.Conf.Redis.MaxActive,
		IdleTimeout: time.Duration(def.Conf.Redis.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", def.Conf.Redis.Addr)
		},
	}
	_, err = pool.Get().Do("ping")
	if err != nil {
		return err
	}

	return nil
}

//etcd init
func InitEtcd() error {
	etcdClient, err = etcd.New(etcd.Config{
		Endpoints:   []string{def.Conf.Etcd.Addr},
		DialTimeout: time.Duration(def.Conf.Etcd.Timeout) * time.Second,
	})
	if err != nil {
		return err
	}

	return nil
}

//close conn, etcd,pool,db (.e.g)
func Close() {
	etcdClient.Close()
	pool.Close()
	db.Close()
}
