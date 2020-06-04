package main

import (
	"github.com/gin-exm/api/def"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers() *gin.Engine {
	r := gin.Default()

	//users handeler
	r.POST("/user", RegisterUser)
	r.POST("/user/login", Login)
	r.GET("/user/:user_name", GetUserInfo)
	r.DELETE("/user/:user_name", Logout)
	r.PUT("/user/:user_name/pwd", ModifyPwd)
	r.PUT("/user/:user_name", ModifyUserInfo)

	return r
}

func main() {
	r := RegisterHandlers()

	//初始化
	if err := initAll(); err != nil {
		panic(err.Error())
	}

	r.Run(":" + def.Conf.Httpport)
}
