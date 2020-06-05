package dbops

import (
	"context"
	"encoding/json"
	"strings"

	mvccpb "github.com/etcd-io/etcd/mvcc/mvccpb"
	"github.com/gin-exm/api/def"
)

func loadProductConfig(key string) error {
	resp, err := etcdClient.Get(context.Background(), key)
	if err != nil {
		return err
	}

	var productInfo []def.ProductConf
	for _, v := range resp.Kvs {
		def.Log.Infoln(v)
		err = json.Unmarshal(v.Value, &productInfo)
		if err != nil {
			return err
		}
	}

	updateProductInfo(productInfo)
	return nil
}

func updateProductInfo(productInfo []def.ProductConf) {
	for _, v := range productInfo {
		def.ProductConfig.Store(v.ProductID, &v)
	}
}

func watchProductKey(key string) {
	def.Log.Infoln("watching key :" + key)
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
