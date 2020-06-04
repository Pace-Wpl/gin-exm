package main

import (
	"github.com/gin-exm/api/dbops"
)

func ValidateUserPwd(uname, pwd string) bool {
	pwdC, err := dbops.GetUserCredential(uname)
	if err != nil || len(pwdC) == 0 || pwd != pwdC {
		return false
	}

	return true
}
