package dbops

import (
	"log"
	"testing"

	"github.com/gin-exm/api/def"
)

func TestMain(m *testing.M) {
	// clearTables()
	m.Run()
	// clearTables()
}

//test user work flow
func TestUserWorkFlow(t *testing.T) {
	t.Run("get user 1\n", testGetUsern)
	t.Run("add user pace\n", testAddUser)
	t.Run("get user pace\n", testGetUser)
}

func testAddUser(t *testing.T) {
	user := &def.User{Username: "pace", Pwd: "wa6602393"}
	if err := AddUser(user); err != nil {
		log.Println(err.Error())
	}
}

func testGetUser(t *testing.T) {
	u := &def.User{}
	if u, err = GetUser("pace"); err != nil {
		log.Println(err.Error())
	}
	log.Println(u)
}

func testGetUsern(t *testing.T) {
	u := &def.User{}
	if u, err = GetUser("55"); err != nil {
		log.Println(err.Error())
	}
	log.Println(u)
}
