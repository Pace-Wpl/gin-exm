package dbops

import (
	"fmt"

	"github.com/gin-exm/api/def"
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

func ListProduct() ([]def.ProductConf, error) {
	var productList []def.ProductConf
	def.ProductConfig.Range(func(key, value interface{}) bool {
		productList = append(productList, value.(def.ProductConf))
		return true
	})

	return productList, nil
}

func GetProduct(pid int) (*def.ProductConf, error) {
	p := &def.ProductConf{}
	v, ok := def.ProductConfig.Load(pid)
	if ok {
		p = v.(*def.ProductConf)
		return p, nil
	}
	return nil, fmt.Errorf("get product id :%d error", pid)
}

func KillProduct(pid int) error {
	return nil
}
