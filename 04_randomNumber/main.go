package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	num := rand.Intn(100) + 1
	fmt.Printf("The random number generated was %d\n", num)
}
