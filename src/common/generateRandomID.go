package common

import (
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano())
}

const chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateShortID(length int) string {
	shortID := make([]byte, length)
	for i := range shortID {
		shortID[i] = chars[rand.Intn(len(chars))]
	}
	return string(shortID)
}