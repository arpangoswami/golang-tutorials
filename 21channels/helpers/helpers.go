package helpers

import (
	"math/rand"
	"time"
)

func RandomIntegerGenerate(n int) int {
	seconds := time.Now().UnixNano()
	rand.Seed(seconds)
	num := rand.Intn(n) + 1
	return num
}
