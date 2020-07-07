package dbops

import (
	"context"
	"encoding/json"
	"strings"

	mvccpb "github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/gin-exm/business/def"
)

var (
	controlHandleChan = make(chan int)
)

func handleControlInfo(controlInfo *def.Control) {
	def.Log.Infoln("handle control info")
	if controlInfo.BusinessSwitch == 0 {
		def.Log.Infoln("close the business mod")
		controlHandleChan <- 0
	}
}

func watchProductKey(key string) {
	def.Log.Debugln("watching key :" + key)
	for {
		rch := etcdClient.Watch(context.Background(), key)
		var controlInfo *def.Control
		var getConfSucc = true

		for wresp := range rch {
			for _, ev := range wresp.Events {
				if ev.Type == mvccpb.DELETE {
					def.Log.Warnln(key + " config deleted")
					continue
				}

				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &controlInfo)
					if err != nil {
						def.Log.Errorln(err.Error())
						getConfSucc = false
						continue
					}
				}
			}

			if getConfSucc {
				def.Log.Infoln("watch product info updata")
				handleControlInfo(controlInfo)
			}
		}

	}
}

func initProductWatcher(key string) {
	go watchProductKey(key)
}

//Etcd任务，包括监听control配置
func PrepareEtcd() error {
	//构造etcd control key，不以 '/' 结尾，加上 '/'
	if strings.HasSuffix(def.Conf.Etcd.PrefixKey, "/") == false {
		def.Conf.Etcd.PrefixKey = def.Conf.Etcd.PrefixKey + "/"
	}
	controlKey := def.Conf.Etcd.PrefixKey + def.Conf.Etcd.ControlKey
	def.Log.Debugln("productKey:" + controlKey)

	//监听etcd配置
	initProductWatcher(controlKey)
	return nil
}
