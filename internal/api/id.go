package api

import (
	"math/rand"
)

const idBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateID() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = idBytes[rand.Intn(len(idBytes))]
	}
	return string(b)
}
