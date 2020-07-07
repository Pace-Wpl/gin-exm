package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/gomodule/redigo/redis"
)

const (
	EtcdKey = "/pace/gin-exm/product"
)

type SecInfoConf struct {
	ProductID int   `json:"product_id"`
	StartTime int64 `json:"start_time"`
	EndTime   int64 `json:"end_time"`
	Status    int   `json:"status"`
	Total     int   `json:"total"`
}

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	defer cli.Close()

	var SecInfoConfArr []SecInfoConf
	SecInfoConfArr = append(
		SecInfoConfArr,
		SecInfoConf{
			ProductID: 1029,
			StartTime: 1594094400,
			EndTime:   1594353600,
			Status:    0,
			Total:     1000,
		},
	)
	SecInfoConfArr = append(
		SecInfoConfArr,
		SecInfoConf{
			ProductID: 1028,
			StartTime: 1594094400,
			EndTime:   1594353600,
			Status:    0,
			Total:     1000,
		},
	)

	data, err := json.Marshal(SecInfoConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//cli.Delete(ctx, EtcdKey)
	//return
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

func setRdist() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println(err.Error())
	}
	_, err = conn.Do("hset", "product_num", "1028", 1000)
	if err != nil {
		log.Println(err.Error())
	}
	_, err = conn.Do("hset", "product_num", "1029", 1000)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	SetLogConfToEtcd()
	// setRdist()
}
