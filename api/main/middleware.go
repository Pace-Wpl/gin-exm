package main

import (
	"net/http"

	"github.com/gin-exm/api/def"
	"github.com/gin-gonic/gin"
)

type connLimiter struct {
	concurrentConn int
	bucket         chan int
}

var (
	connLi = &connLimiter{
		concurrentConn: def.Conf.StreamLimit,
		bucket:         make(chan int, def.Conf.StreamLimit),
	}
)

func StreamLimitdMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//判断请求连接数是否超过规定连接数
		if !connLi.getConn() {
			c.JSON(http.StatusOK, def.ErrorServerBusy)
			return
		}

		c.Next()

		//处理结束，释放连接
		connLi.releaseConn()
	}
}

func (cl *connLimiter) getConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		def.Log.Warnln("Reached the rate limitation.")
		return false
	}

	cl.bucket <- 1
	def.Log.Infoln("get a connect")
	return true
}

func (cl *connLimiter) releaseConn() {
	_ = <-cl.bucket
	def.Log.Infoln("release a connect")
}
