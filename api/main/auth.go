package main

import (
	"github.com/gin-exm/api/dbops"
	"github.com/gin-exm/api/def"
	"github.com/gin-exm/api/session"
)

func ValidateUserPwd(uname, pwd string) bool {
	pwdC, err := dbops.GetUserCredential(uname)
	if err != nil || len(pwdC) == 0 || pwd != pwdC {
		return false
	}

	return true
}

func ValidateToken(token *def.Session) bool {
	uID, ok := session.IsSessionExpried(token.ID)
	if !ok {
		if uID != token.UserId {
			def.Log.Warnln("token user id not equal to token's user id")
			return false
		}
		return true
	}
	return false
}
