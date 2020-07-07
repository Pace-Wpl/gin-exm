package dbops

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/gin-exm/api/def"
	"github.com/gomodule/redigo/redis"
)

/*
 ETCD SECTION
*/

func loadProductConfig(key string) error {
	resp, err := etcdClient.Get(context.Background(), key)
	if err != nil {
		return err
	}

	var productInfo []def.ProductConf
	for _, v := range resp.Kvs {
		err = json.Unmarshal(v.Value, &productInfo)
		if err != nil {
			return err
		}
	}

	def.Log.Infoln("load product config")
	updateProductInfo(productInfo)
	return nil
}

func updateProductInfo(productInfo []def.ProductConf) {
	def.Log.Infoln("updata product info")
	for _, v := range productInfo {
		t := v
		def.Log.Infoln(v)
		productConfigMap.Store(v.ProductID, &t)
	}
}

func watchProductKey(key string) {
	def.Log.Debugln("watching key :" + key)
	for {
		rch := etcdClient.Watch(context.Background(), key)
		var productInfo []def.ProductConf
		var getConfSucc = true

		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					def.Log.Warnln(key + " config deleted")
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &productInfo)
					if err != nil {
						def.Log.Errorln(err.Error())
						getConfSucc = false
						continue
					}
				}
			}

			if getConfSucc {
				def.Log.Infoln("watch product info updata")
				updateProductInfo(productInfo)
			}
		}

	}
}

func initProductWatcher(key string) {
	go watchProductKey(key)
}

//Etcd任务，包括初次获取product配置，监听product配置
func PrepareEtcd() error {
	//构造etcd product key，不以 '/' 结尾，加上 '/'
	if strings.HasSuffix(def.Conf.Etcd.PrefixKey, "/") == false {
		def.Conf.Etcd.PrefixKey = def.Conf.Etcd.PrefixKey + "/"
	}
	productKey := def.Conf.Etcd.PrefixKey + def.Conf.Etcd.ProductKey
	def.Log.Infoln("productKey:" + productKey)

	//初次获取etcd配置
	if err := loadProductConfig(productKey); err != nil {
		return err
	}

	//监听etcd配置
	initProductWatcher(productKey)
	return nil
}

/*
 REDIS SECTION
*/

func writeHandle(key string) {
	for {
		select {
		case req := <-reqChan:
			conn := pool.Get()

			data, err := json.Marshal(req)
			if err != nil {
				def.Log.Errorf("json marshal failed, err:%v", err)
				conn.Close()
				continue
			}

			_, err = conn.Do("LPUSH", key, string(data))
			if err != nil {
				def.Log.Errorf("lpush failed, err:%v, req:%v", err, req)
				conn.Close()
				continue
			}

			conn.Close()
		}
	}
}

func readHandle(key string) {
	for {
		conn := pool.Get()

		reply, err := redis.Values(conn.Do("BRPOP", key, 0))

		if err != nil || len(reply) != 2 {
			def.Log.Error("pop from queue failed, err:%v", err)
			break
		}

		data, ok := reply[1].([]byte)
		if !ok {
			def.Log.Error("pop from queue failed, err:%v", err)
			continue
		}

		var result *def.ResultSecKill
		err = json.Unmarshal([]byte(data), &result)
		if err != nil {
			def.Log.Errorf("json.Unmarshal failed, err:%v", err)
			conn.Close()
			continue
		}

		userKey := fmt.Sprintf("%s-%d", result.UserId, result.ProductId)

		resultChan, ok := userConnMap.Load(userKey)
		if !ok {
			conn.Close()
			def.Log.Warnf("user not found:%v", userKey)
			continue
		}

		resultChan.(chan *def.ResultSecKill) <- result
		conn.Close()
	}
}

func redisListen() {
	go writeHandle(def.Conf.Redis.SecReqQueue)
	go readHandle(def.Conf.Redis.SecRespQueue)
}

//Redis监听任务
func PrepareRedis() error {
	//redis监听任务
	redisListen()
	return nil
}
