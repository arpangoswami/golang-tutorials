package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter marks obtained: ")
	reader := bufio.NewReader(os.Stdin)
	grade, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	grade = strings.TrimSpace(grade)
	marks, errMar := strconv.ParseFloat(grade, 64)
	if errMar != nil {
		log.Fatal(errMar)
	}
	passMarks := 30
	goodMarks := 60
	if marks >= float64(goodMarks) {
		fmt.Printf("You obtained %f marks which was better than good i.e. %f\n", marks, float64(goodMarks))
	} else if marks >= float64(passMarks) {
		fmt.Printf("You obtained %f marks which was better than passing marks i.e. %f but failed to achieve good grade as your marks were lower than %f\n", marks, float64(goodMarks), float64(passMarks))
	} else {
		fmt.Printf("You failed as you got less than %f", float64(passMarks))
	}
}
