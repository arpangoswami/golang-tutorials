package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	seconds := time.Now().Unix()
	rand.Seed(seconds)
	originalNum := rand.Intn(100) + 1
	fmt.Println("Guess the number between 1 and 100")
	fmt.Println("You get 7 guesses to guess the number")
	reader := bufio.NewReader(os.Stdin)
	correct := false
	for guess := 0; guess < 7; guess++ {
		fmt.Println("You have", 7-guess, "guesses left")

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if number == originalNum {
			correct = true
			fmt.Println("You guessed correctly", originalNum, "is the correct number.")
			break
		} else if number < originalNum {
			fmt.Println("Your guess was a lower than the original number")
		} else {
			fmt.Println("Your guess was a bit higher than the original number")
		}
	}
	if !correct {
		fmt.Println("You ran out of guesses!! The original number was", originalNum)
	}
}
