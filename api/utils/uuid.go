package utils

import (
	"crypto/rand"
	"fmt"
	"io"
	"strconv"
	"time"
)

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits;
	uuid[8] = uuid[4]&^0xc0 | 0x80
	// version 4 (pseudo-random);
	uuid[6] = uuid[3]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x", uuid[0:4], uuid[4:8], uuid[8:12], uuid[12:16]), nil
}

func NewNance() (string, error) {
	uuid := make([]byte, 4)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", uuid[0:4]), nil
}

func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}
