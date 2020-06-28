package dbops

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/gin-exm/business/def"
)

func HandleSecKill(req *def.ReqSecKill) (*def.ResultSecKill, error) {
	data := []byte(strconv.Itoa(req.ProductID) + "_" + def.Conf.CryptoStr + "_" + req.UserID + "_" + req.Time)
	h := md5.New()
	h.Write(data)
	result := &def.ResultSecKill{
		ProductId: req.ProductID, UserId: req.UserID, Mes: "秒杀成功",
		Token: hex.EncodeToString(h.Sum(nil)),
	}

	return result, nil
}
