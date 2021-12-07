package main

import (
	"errors"
	"fmt"
)

func divide(dividend int, divisor int) (int, int, error) {
	if divisor == 0 {
		err := errors.New("division by zero called not possible")
		return 0, 0, err
	}
	return dividend / divisor, dividend % divisor, nil
}
func main() {
	var dividend int
	var divisor int
	fmt.Println("Enter two integers dividend and divisor")
	fmt.Scanf("%d %d", &dividend, &divisor)
	quotient, remainder, err := divide(dividend, divisor)
	if err != nil {
		fmt.Println("Error encountered: ", err)
	} else {
		fmt.Printf("After dividing %d with %d we get the quotient as %d and remainder as %d\n", dividend, divisor, quotient, remainder)
	}
}
