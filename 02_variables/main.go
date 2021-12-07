package main

import "fmt"

func main() {
	var originalCount int = 10
	var eatenCount int = 4
	var length, width float64 = 2.4, 4.5
	fmt.Printf("I started with %d apples.\n", originalCount)
	fmt.Printf("Some jerk ate %d apples.\n", eatenCount)
	fmt.Printf("There are now %d apples left.\n", originalCount-eatenCount)
	fmt.Printf("%f %f\n", length, width)
}
