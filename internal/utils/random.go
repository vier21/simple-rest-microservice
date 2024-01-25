package utils

import (
	"math/rand"
	"time"
)

var (
	char = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTU"
	rnd  = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandomString(length int) string {
	b := make([]byte, length)

	for i := range b {
		b[i] = char[rnd.Intn(len(char))]
	}
	return string(b)
}
