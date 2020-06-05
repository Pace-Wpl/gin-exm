package main

import (
	"net/http"

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

	uu := &def.User{Username: u.Name, Pwd: u.Password, Icon: def.DEFAULT_ICON}
	if err := dbops.AddUser(uu); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	sID, err := session.GenerateNewSession(uu.Username)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.SetCookie(def.Conf.CookieKey, sID, def.Conf.SessionExpired, "/", def.Conf.Domain, false, true)
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

	sID, err := session.GenerateNewSession(u.Name)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.SetCookie(def.Conf.CookieKey, sID, def.Conf.SessionExpired, "/", def.Conf.Domain, false, true)
	c.JSON(http.StatusOK, &def.RespMes{Mes: "welcome " + u.Name, Code: 200})
}

func GetUserInfo(c *gin.Context) {
	uname := c.Param("user_name")
	u, err := dbops.GetUser(uname)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
	}

	c.JSON(http.StatusOK, u)
}

func Logout(c *gin.Context) {
	cookie, err := c.Request.Cookie(def.Conf.CookieKey)
	if err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
	}

	if err = session.DelSession(cookie.Value); err != nil {
		def.Log.Warnln(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
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
