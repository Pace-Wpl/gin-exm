package session

import (
	"github.com/gin-exm/api/dbops"
	"github.com/gin-exm/api/def"
	"github.com/gin-exm/api/utils"
)

func GenerateNewSession(username string) (string, error) {
	id, _ := utils.NewUUID()
	session := &def.Session{ID: id, Name: username}

	if err := dbops.SetSession(session); err != nil {
		return "", err
	}
	return id, nil
}

func ReSetSession(sid, uname string) error {
	session := &def.Session{ID: sid, Name: uname}

	if err := dbops.SetSession(session); err != nil {
		return err
	}
	return nil
}

func DelSession(sid string) error {
	return dbops.DeleteSession(sid)
}

func IsSession(sid string) (string, bool) {
	return dbops.IsSessionExpried(sid)
}
