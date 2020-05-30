package dbops

import (
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
