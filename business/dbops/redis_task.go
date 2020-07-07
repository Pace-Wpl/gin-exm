package dbops

import (
	"context"
	"encoding/json"
	"time"

	"github.com/gin-exm/business/def"
	"github.com/gomodule/redigo/redis"
)

var (
	readHandleChan  = make(chan *def.ReqSecKill, def.Conf.ReqChannelBuffer)
	writeHandleChan = make(chan *def.ResultSecKill, def.Conf.RespChannelBuffer)
)

//监听request MQ，并写入 read handle channel
func handleReader(ctx context.Context, key string) {
	def.Log.Info("read goroutine running")
	for {
		conn := pool.Get()
		for {
			ret, err := redis.Values(conn.Do("BRPOP", key))
			if err != nil || len(ret) != 2 {
				def.Log.Error("pop from queue failed, err:%v", err)
				break
			}

			data, ok := ret[1].([]byte)
			if !ok {
				def.Log.Error("pop from queue failed, err:%v", err)
				continue
			}

			var req *def.ReqSecKill
			err = json.Unmarshal([]byte(data), &req)
			if err != nil {
				def.Log.Error("unmarshal to secrequest failed, err:%v", err)
				continue
			}

			now := time.Now().Unix()
			if now-req.AccessTime.Unix() >= int64(def.Conf.RequestWaitTimeOut) {
				def.Log.Warn("req[%v] is expire", req)
				continue
			}

			timer := time.NewTicker(time.Duration(def.Conf.ResponseSendTimeOut) * time.Second)
			select {
			case <-ctx.Done():
				def.Log.Warn("goroutine 'handleReader' closing...")
				timer.Stop()
				return
			case readHandleChan <- req:
			case <-timer.C:
				def.Log.Warn("send to handle chan timeout, req:%v", req)
				break
			}
			timer.Stop()
		}
		conn.Close()
	}
}

//监听write handel channel获取response 写入 response MQ
func handleWrite(ctx context.Context, key string) {
	def.Log.Debug("handle write running")

	for {
		select {
		case <-ctx.Done():
			def.Log.Warn("goroutine 'handleWrite' closing...")
			return
		case res := <-writeHandleChan:
			err := sendToRedis(res, key)
			if err != nil {
				def.Log.Error("send to redis, err:%v, res:%v", err, res)
			}
		}
	}
}

func sendToRedis(res *def.ResultSecKill, key string) error {

	data, err := json.Marshal(res)
	if err != nil {
		def.Log.Error("marshal failed, err:%v", err)
		return err
	}

	conn := pool.Get()
	defer conn.Close()
	_, err = conn.Do("RPUSH", key, string(data))
	if err != nil {
		def.Log.Warn("rpush to redis failed, err:%v", err)
		return err
	}

	return nil
}

//接受request MQ，处理，response MQ
func handleBusiness(ctx context.Context) {

	def.Log.Info("handle user running")
	for req := range readHandleChan {
		def.Log.Info("begin process request:%v", req)
		res, err := HandleSecKill(req)
		if err != nil {
			def.Log.Warn("process request %v failed, err:%v", err)
			res = &def.ResultSecKill{
				ProductId: req.ProductID, UserId: req.UserID, Mes: err.Error(), Token: "",
			}
		}

		timer := time.NewTicker(time.Duration(def.Conf.ResponseSendTimeOut) * time.Second)
		select {
		case <-ctx.Done():
			def.Log.Warn("goroutine 'handleBusiness' closing...")
			timer.Stop()
			return
		case writeHandleChan <- res:
		case <-timer.C:
			def.Log.Warn("send to response chan timeout, res:%v", res)
			break
		}
		timer.Stop()
	}
	return
}

//开启redis goroutine，并等待任务
func redisTask() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < def.Conf.ReadGoroutineNum; i++ {
		go handleReader(ctx, def.Conf.Redis.SecReqQueue)
	}
	for i := 0; i < def.Conf.WriteGoroutineNum; i++ {
		go handleWrite(ctx, def.Conf.Redis.SecRespQueue)
	}
	for i := 0; i < def.Conf.HandleGoroutineNum; i++ {
		go handleBusiness(ctx)
	}

	def.Log.Debug("all process goroutine started")
	select {
	case <-controlHandleChan:
		cancel()
		def.Log.Debug("close all goroutine")
	}
	return
}

//Redis任务，包括接受请求，处理请求，响应请求
func PrepareRedis() error {
	//redis任务
	go redisTask()
	return nil
}
