package util

import (
	"math/rand"
	"time"
)

func GetRand(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
