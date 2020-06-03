package main

import (
	"log"
	"net/http"

	"github.com/gin-exm/api/dbops"
	"github.com/gin-exm/api/def"
	"github.com/gin-exm/api/session"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	u := &def.ReqUser{}
	if err := c.BindJSON(u); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}

	uu := &def.User{Username: u.Name, Pwd: u.Password, Icon: def.DEFAULT_ICON}
	if err := dbops.AddUser(uu); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	sID, err := session.GenerateNewSession(uu.Username)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.SetCookie("sessionID", sID, def.SESSION_EXPIRED, "/", "localhost", false, true)
	c.JSON(http.StatusCreated, u)
}

func Login(c *gin.Context) {
	u := &def.ReqUser{}
	if err := c.BindJSON(u); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorRequestBodyPaseFailed)
		return
	}

	pwd, err := dbops.GetUserCredential(u.Name)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, def.ErrorInternalError)
	}

	if pwd != u.Password {
		c.JSON(http.StatusUnauthorized, def.ErrorNotAuthUser)
	}

	sID, err := session.GenerateNewSession(u.Name)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, def.ErrorInternalError)
		return
	}

	c.SetCookie("sessionID", sID, def.SESSION_EXPIRED, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"username": u.Name})
}

func GetUserInfo(c *gin.Context) {
	uname := c.Param("user_name")
	u := &def.ReqUser{Name: uname, Password: "123"}
	c.JSON(200, u)
	// //验证用户是否登陆
	// if !ValidateLogin(w, r) {
	// 	log.Printf("Unauthorized user\n")
	// 	return
	// }

	// uname := p.ByName("user_name")
	// user, err := dbops.GetUser(uname)
	// if err != nil {
	// 	sendErrorResponse(w, def.ErrorDBError)
	// 	return
	// }

	// userInfo := &def.UserInfo{Id: user.Id, Pwd: user.Pwd, Name: user.Username}
	// if resp, err := json.Marshal(userInfo); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// } else {
	// 	sendNormalResponse(w, string(resp), 200)
	// }

}

func Logout(c *gin.Context) {
	// sid := r.Header.Get(HEADER_FIELD_SESSION)
	// if len(sid) == 0 {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// }
	// session.DeleteExpiredSession(sid)
	// sendNormalResponse(w, "Logout ok !", 200)
	// //	io.WriteString(w, "user logout!")
	resp := &def.RespMes{Mes: "logout", Code: 200}
	c.JSON(200, resp)
}

func ModifyPwd(c *gin.Context) {
	// //验证用户
	// if !ValidateUser(w, r, p) {
	// 	log.Printf("Unauthorized user\n")
	// 	return
	// }

	// uname := p.ByName("user_name")
	// res, _ := ioutil.ReadAll(r.Body)
	// ubody := &def.UserModifyPwd{}
	// if err := json.Unmarshal(res, ubody); err != nil {
	// 	log.Printf("unmarshal error!")
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// }

	// if !ValidateUserPwd(w, ubody.PTPwd, uname) {
	// 	log.Printf("pass word error!")
	// 	return
	// }

	// if err := dbops.ModifyUserPwd(uname, ubody.CPwd); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// }

	// if resp, err := json.Marshal(ubody); err != nil {
	// 	log.Printf("marshal error!")
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// } else {
	// 	sendNormalResponse(w, string(resp), 200)
	// }

	// //	io.WriteString(w, "Modify password!")
	uname := c.Param("user_name")
	resp := &def.RespMes{Mes: uname, Code: 200}
	c.JSON(200, resp)
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
