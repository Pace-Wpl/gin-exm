package dbops

import (
	"encoding/json"
	"errors"

	"github.com/gin-exm/api/def"
	"github.com/gomodule/redigo/redis"
)

func SetSession(session *def.Session) error {
	conn := pool.Get()
	if conn == nil {
		return errors.New("get redis connection error")
	}
	defer conn.Close()

	value, err := json.Marshal(session)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", session.ID, value, "EX", def.Conf.SessionExpired)
	if err != nil {
		return err
	}
	return nil
}

func IsSessionExpried(sid string) (string, bool) {
	conn := pool.Get()
	if conn == nil {
		return "", true
	}
	defer conn.Close()

	ok, _ := redis.Bool(conn.Do("EXISTS", sid))
	if ok {
		res, err := redis.Bytes(conn.Do("GET", sid))
		if err != nil {
			def.Log.Warnln(err.Error())
			return "", true
		}

		session := &def.Session{}
		if err = json.Unmarshal(res, session); err != nil {
			def.Log.Warnln(err.Error())
			return "", true
		}
		return session.UserId, false
	}
	return "", true
}

func DeleteSession(sid string) error {
	conn := pool.Get()
	if conn == nil {
		return errors.New("get redis connection error")
	}
	defer conn.Close()

	_, err := conn.Do("DEL", sid)
	if err != nil {
		return err
	}
	return nil
}
