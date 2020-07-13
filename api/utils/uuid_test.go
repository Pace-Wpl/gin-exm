package utils

import (
	"fmt"
	"testing"
)

func TestUuid(t *testing.T) {
	t.Run("...", testNance)

}

func testUuid(t *testing.T) {
	id, _ := NewUUID()
	fmt.Printf("id:%v", id)
}

func testNance(t *testing.T) {
	id, _ := NewNance()
	fmt.Printf("id:%v", id)
}
