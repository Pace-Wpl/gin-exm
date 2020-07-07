package dbops

import (
	"time"

	etcd "github.com/etcd-io/etcd/clientv3"
	"github.com/gin-exm/business/def"
	"github.com/gomodule/redigo/redis"
)

var (
	etcdClient *etcd.Client
	pool       *redis.Pool
	err        error
)

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
	close(readHandleChan)
	close(writeHandleChan)
	close(controlHandleChan)
}
