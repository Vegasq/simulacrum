package utils

import (
	"crypto/rand"
	"fmt"
	"os"
)

// GenUUID generate random UUID
func GenUUID() string {
	f, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// GenMAC generate random MAC
func GenMAC() string {
	buf := make([]byte, 3)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	buf[0] |= 2
	return fmt.Sprintf("52:54:00:%02x:%02x:%02x", buf[0], buf[1], buf[2])
}
