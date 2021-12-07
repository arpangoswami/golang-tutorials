package main

import (
	"errors"
	"log"
)

func Divide(x float32, y float32) (float32, error) {
	var result float32
	if y == 0.0 {
		return result, errors.New("division by zero not possible")
	}
	result = x / y
	return result, nil
}

func main() {
	result, err := Divide(100.0, 0.0)
	if err != nil {
		log.Println("Error: ", err)
		return
	}
	log.Println("Result is:", result)
}
