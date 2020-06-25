package dbops

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-exm/api/def"
)

var (
	productConfigMap sync.Map
	userConnMap      sync.Map
	reqChan          = make(chan *def.ReqSecKill, def.Conf.ReqChannelBuffer)
)

func AddUser(user *def.User) error {
	if err := db.Table("users").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUser(uname string) (*def.User, error) {
	user := &def.User{}
	if err := db.Table("users").Where("username = ?", uname).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserCredential(uname string) (string, error) {
	user := &def.User{}
	if err := db.Table("users").Where("username = ?", uname).Select("pwd").Find(user).Error; err != nil {
		return "", err
	}

	return user.Pwd, nil
}

func ModifyPwd(uname, pwd string) error {
	if err := db.Table("users").Where("username = ?", uname).Update("pwd", pwd).Error; err != nil {
		return err
	}
	return nil
}

func ListProduct() ([]*def.ProductConf, error) {
	var productList []*def.ProductConf
	productConfigMap.Range(func(key, value interface{}) bool {
		productList = append(productList, value.(*def.ProductConf))
		return true
	})

	return productList, nil
}

func getProduct(pid int) (*def.ProductConf, error) {
	p := &def.ProductConf{}
	v, ok := productConfigMap.Load(pid)
	if ok {
		p = v.(*def.ProductConf)
		return p, nil
	}
	return nil, fmt.Errorf("get product id :%d error", pid)
}

func ObtainProductInfo(pid int) (*def.RespProductInfo, error) {
	p, err := getProduct(pid)
	if err != nil {
		return nil, err
	}

	pp := &def.RespProductInfo{ProductID: p.ProductID, StartTime: p.StartTime, EndTime: p.EndTime,
		Status: p.Status, Activity: 0, Total: p.Total}
	//活动开始,时间，状态判断,时间未到，标记活动未开始，时间>start.time  <end.time 并且 status 为 nomal,标记活动开始
	//其他:标记活动结束
	if time.Now().Unix() < pp.StartTime {
		pp.Activity = def.PRODUCT_ACTIVITY_PRE
	} else if time.Now().Unix() >= pp.StartTime && time.Now().Unix() <= pp.EndTime && pp.Status == def.PRODUCT_STATUS_NOMAL {
		pp.Activity = def.PRODUCT_ACTIVITY_BEGIN
	} else {
		pp.Activity = def.PRODUCT_ACTIVITY_END
	}

	return pp, nil
}

func KillProduct(req *def.ReqSecKill) (*def.ResultSecKill, error) {
	ch := make(chan *def.ResultSecKill)
	key := fmt.Sprintf("%s-%d", req.UserID, req.ProductID)
	userConnMap.Store(key, ch)

	reqChan <- req

	//过期时间
	ticker := time.NewTicker(time.Duration(def.Conf.ReqTimeout) * time.Second)

	defer func() {
		ticker.Stop()
		userConnMap.Delete(key)
	}()

	select {
	case <-ticker.C:
		return nil, errors.New("time out")
	case result := <-ch:
		return result, nil
	}
}
