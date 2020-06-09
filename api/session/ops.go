package session

import (
	"github.com/gin-exm/api/dbops"
	"github.com/gin-exm/api/def"
	"github.com/gin-exm/api/utils"
)

func GenerateNewSession(username string) (*def.Session, error) {
	id, _ := utils.NewUUID()
	session := &def.Session{ID: id, UserId: username}

	if err := dbops.SetSession(session); err != nil {
		return nil, err
	}
	return session, nil
}

func ReSetSession(sid, uname string) error {
	session := &def.Session{ID: sid, UserId: uname}

	if err := dbops.SetSession(session); err != nil {
		return err
	}
	return nil
}

func DelSession(sid string) error {
	return dbops.DeleteSession(sid)
}

func IsSessionExpried(sid string) (string, bool) {
	return dbops.IsSessionExpried(sid)
}
