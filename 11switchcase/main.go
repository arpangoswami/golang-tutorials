package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	switch diceNumber {
	case 1:
		fmt.Println("You can open your account. If already opened Move up two spots")
	case 2:
		fmt.Println("You can move up two spots")
	case 3:
		fmt.Println("You can move up three spots")
	case 4:
		fmt.Println("You can move up four spots")
	case 5:
		fmt.Println("You can move up five spots")
	case 6:
		fmt.Println("You can move up six spots. You can now take another turn")
	default:
		fmt.Println("What was that??")
	}
}
