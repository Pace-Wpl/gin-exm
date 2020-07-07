package main

import (
	"io"
	"os"

	"github.com/gin-exm/business/dbops"
	"github.com/gin-exm/business/def"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

func initConfig() error {
	if err := configor.Load(def.Conf, def.CONFIG_DIR); err != nil {
		return err
	}
	log.Infoln(def.Conf)
	return nil
}

func initLog() error {
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

func initDB() error {
	var err error
	if err = dbops.InitRedis(); err != nil {
		return err
	}

	if err = dbops.InitEtcd(); err != nil {
		return err
	}

	return nil
}

//初始化，包括配置文件，日志，redis，etcd模块
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

	//初始化db：redis, etcd
	if err = initDB(); err != nil {
		def.Log.Fatal(err.Error())
		return err
	}

	return err
}

//开启任务,Redis任务
func startAll() error {
	var err error

	if err = dbops.PrepareEtcd(); err != nil {
		def.Log.Fatal(err.Error())
		return err
	}

	if err = dbops.PrepareRedis(); err != nil {
		def.Log.Fatal(err.Error())
		return err
	}

	return nil
}

func close() {
	dbops.Close()
}
