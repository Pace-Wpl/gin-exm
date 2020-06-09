package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-exm/api/dbops"
	"github.com/gin-exm/api/def"
	"github.com/gin-exm/api/session"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	u := &def.ReqUser{}
	if err := c.BindJSON(u); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}

	uu := &def.User{Username: u.Name, Pwd: u.Password, Icon: def.Conf.DefaultIcon}
	if err := dbops.AddUser(uu); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	token, err := session.GenerateNewSession(uu.Username)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.SetCookie(def.Conf.CookieKey1, token.ID, def.Conf.SessionExpired, "/", def.Conf.Domain, false, true)
	c.SetCookie(def.Conf.CookieKey2, token.UserId, def.Conf.SessionExpired, "/", def.Conf.Domain, false, true)
	c.JSON(http.StatusCreated, &def.RespMes{Mes: "welcome " + u.Name, Code: 200})
}

func Login(c *gin.Context) {
	u := &def.ReqUser{}
	if err := c.BindJSON(u); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}

	//密码验证
	if !ValidateUserPwd(u.Name, u.Password) {
		c.JSON(http.StatusUnauthorized, def.ErrorNotAuthUser)
		return
	}

	token, err := session.GenerateNewSession(u.Name)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.SetCookie(def.Conf.CookieKey1, token.ID, def.Conf.SessionExpired, "/", def.Conf.Domain, false, true)
	c.SetCookie(def.Conf.CookieKey2, token.UserId, def.Conf.SessionExpired, "/", def.Conf.Domain, false, true)
	c.JSON(http.StatusOK, &def.RespMes{Mes: "welcome " + u.Name, Code: 200})
}

func GetUserInfo(c *gin.Context) {
	uname := c.Param("user_name")
	u, err := dbops.GetUser(uname)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.JSON(http.StatusOK, u)
}

func Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie(def.Conf.CookieKey1)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusOK, def.ErrorUserNotLogin)
		return
	}

	if err = session.DelSession(cookie.Value); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.JSON(http.StatusOK, &def.RespMes{Mes: "logout successful!", Code: 200})
}

func ModifyPwd(c *gin.Context) {
	uname := c.Param("user_name")
	u := &def.ReqModifyPwd{}
	if err := c.BindJSON(u); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}

	//密码验证
	if !ValidateUserPwd(uname, u.Pwd) {
		c.JSON(http.StatusUnauthorized, def.ErrorNotAuthUser)
		return
	}

	if err := dbops.ModifyPwd(uname, u.NewPwd); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.JSON(http.StatusOK, &def.RespMes{Mes: "modify successful!", Code: 200})
}

func ModifyUserInfo(c *gin.Context) {
	// if !ValidateUser(w, r) {
	// 	log.Printf("Unauthorized user\n")
	// 	return
	// }

	// uname := p.ByName("user_name")
	// res, _ := ioutil.ReadAll(r.Body)
	// ubody := &def.UserModifyInfo{}
	// if err := json.Unmarshal(res, ubody); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// }

	// if err := dbops.ModifyUserInfo(uname, ubody.CPwd); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// }

	// if resp, err := json.Marshal(ubody); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// } else {
	// 	sendNormalResponse(w, string(resp), 200)
	// }
	uname := c.Param("user_name")
	resp := &def.RespMes{Mes: uname, Code: 200}
	c.JSON(200, resp)
}

func ListProduct(c *gin.Context) {
	var productList []*def.ProductConf
	var err error
	if productList, err = dbops.ListProduct(); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.JSON(http.StatusOK, productList)
}

func GetProduct(c *gin.Context) {
	pid, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}

	p, err := dbops.GetProduct(pid)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusOK, &def.RespMes{Mes: "product sell out or not exist!", Code: 500})
		return
	}

	c.JSON(http.StatusOK, p)
}

func ProductSecKill(c *gin.Context) {
	cookie1, err := c.Request.Cookie(def.Conf.CookieKey1)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusOK, def.ErrorUserNotLogin)
		return
	}
	cookie2, err := c.Request.Cookie(def.Conf.CookieKey2)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusOK, def.ErrorUserNotLogin)
		return
	}
	token := &def.Session{ID: cookie1.Value, UserId: cookie2.Value}
	//用户token验证
	if ValidateToken(token) {
		c.JSON(http.StatusOK, def.ErrorUserNotLogin)
		return
	}

	pid, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}
	s := c.Query("src")
	a := c.Query("authcode")
	t := c.Query("time")
	n := c.Query("nance")
	//cookie 获取uid
	uid := token.UserId

	ReqKill := &def.ReqSecKill{Sourct: s, ProductID: pid, AuthCode: a, Time: t, Nance: n, AccessTime: time.Now(), UserID: uid}

	//用户访问控制
	if AntiSpam(ReqKill) {
		def.Log.Warnln("用户:" + ReqKill.UserID + "访问过多！")
		c.JSON(http.StatusOK, &def.ErrorServerBusy)
		return
	}
	c.JSON(200, gin.H{"message": "秒杀访问成功"})
}
