package main

import (
	"io"
	"os"

	"github.com/gin-exm/api/dbops"
	"github.com/gin-exm/api/def"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

func initConfig() error {
	if err := configor.Load(def.Conf, def.CONFIG_DIR); err != nil {
		return err
	}
	log.Info(def.Conf)
	return nil
}

func initLog() error {
	// init gin log
	gin.DisableConsoleColor()
	f, err := os.OpenFile(def.Conf.Log.GinLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//init sys log
	f1, err := os.OpenFile(def.Conf.Log.SysLogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}
	def.Log.Out = io.MultiWriter(f1, os.Stdout)

	def.Log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return nil
}

func initEtcd() error {
	return nil
}

func initDB() error {
	return dbops.InitDB()
}

func initAll() error {
	var err error

	//初始化配置文件
	if err = initConfig(); err != nil {
		log.Fatal(err.Error())
		return err
	}

	//初始化日志
	if err = initLog(); err != nil {
		log.Fatal(err.Error())
		return err
	}

	//初始化etcd
	if err = initEtcd(); err != nil {
		def.Log.Fatal(err.Error())
		return err
	}

	//初始化db：mysql，redis
	if err = initDB(); err != nil {
		def.Log.Fatal(err.Error())
		return err
	}

	return err
}
