package main

import (
	"github.com/gin-exm/api/def"
	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	u := &def.ReqUser{}
	if err := c.BindJSON(u); err != nil {
		c.JSON(400, def.ErrorRequestBodyPaseFailed)
		return
	}
	c.JSON(200, u)
}

func Login(c *gin.Context) {
	uname := c.PostForm("name")
	ps := c.PostForm("password")
	if uname == "pace" && ps == "123" {
		resp := &def.RespMes{Code: 002, Mes: "登录成功！"}
		c.JSON(200, resp)
	}
	c.JSON(401, def.ErrorNotAuthUser)
	// res, _ := ioutil.ReadAll(r.Body)
	// ubody := &def.UserCredential{}

	// if err := json.Unmarshal(res, ubody); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// }

	// uname := p.ByName("user_name")
	// log.Printf("url name: %v", uname)
	// log.Printf("request name : %v", ubody.Username)
	// if uname != ubody.Username {
	// 	sendErrorResponse(w, def.ErrorNotAuthUser)
	// 	return
	// }

	// if !ValidateUserPwd(w, ubody.Pwd, ubody.Username) {
	// 	log.Printf("pass word error!")
	// 	return
	// }

	// id := session.GenerateNewSessionId(ubody.Username)
	// sup := &def.SignedUp{Success: true, SessionId: id}

	// if resp, err := json.Marshal(sup); err != nil {
	// 	sendErrorResponse(w, def.ErrorInternalFaults)
	// 	return
	// } else {
	// 	sendNormalResponse(w, string(resp), 201)
	// }
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
