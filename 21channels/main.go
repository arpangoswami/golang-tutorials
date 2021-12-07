package main

import (
	"channels/helpers"
	"fmt"
)

const numPool = 10

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomIntegerGenerate(numPool)
	intChan <- randomNumber
}
func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	fmt.Println(num)
}
