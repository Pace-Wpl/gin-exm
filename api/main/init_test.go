package main

import (
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestInitWorkFlow(t *testing.T) {
	t.Run("init\n", testInit)
}

func testInit(t *testing.T) {
	if err := initAll(); err != nil {
		log.Println(err.Error())
	}
}
