package main

import "fmt"

func swap(oneptr *int, twoptr *int) {
	x := *oneptr
	*oneptr = *twoptr
	*twoptr = x
}
func main() {
	fmt.Println("Welcome to class of pointers")
	one := 1
	oneptr := &one
	fmt.Println("Value of one is: ", one)
	fmt.Println("Address of one is: ", oneptr)
	var two int = 2
	var twoPtr *int = &two
	fmt.Println("Value of one is: ", two)
	fmt.Println("Value of one via ptr is: ", *twoPtr)
	swap(oneptr, twoPtr)
	fmt.Println("Value of one variable after swap is: ", one)
	fmt.Println("Value of two variable after swap is: ", two)
}
