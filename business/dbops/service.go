package dbops

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"

	"github.com/gin-exm/business/def"
	"github.com/gomodule/redigo/redis"
)

func HandleSecKill(req *def.ReqSecKill) (*def.ResultSecKill, error) {
	if err := reduceStock(req.ProductID, -1); err != nil {
		return nil, err
	}
	data := []byte(strconv.Itoa(req.ProductID) + "_" + def.Conf.CryptoStr + "_" + req.UserID + "_" + req.Time)
	h := md5.New()
	h.Write(data)
	result := &def.ResultSecKill{
		ProductId: req.ProductID, UserId: req.UserID, Mes: "秒杀成功",
		Token: hex.EncodeToString(h.Sum(nil)), Nance: req.Nance,
	}

	return result, nil
}

//redis 增加库存num
func reduceStock(pid, num int) error {
	pidStr := strconv.Itoa(pid)
	conn := pool.Get()
	defer conn.Close()

	stock, err := redis.Int(conn.Do("hget", "product_num", pidStr))
	if stock <= 0 || err != nil {
		return errors.New("sell out")
	}

	_, err = conn.Do("hincrby", "product_num", pidStr, num)
	if err != nil {
		return err
	}
	return nil
}

// //redis 库存增加1
// func backStock(pid int) error {
// 	pidStr := strconv.Itoa(pid)
// 	conn := pool.Get()
// 	_, err := conn.Do("incr", pidStr)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
