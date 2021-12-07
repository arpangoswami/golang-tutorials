package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Welcome to file handling!")
	content := "This will go inside the file, Arpan Goswami creator. YOLOOO"
	file, err := os.Create("./sample_file.txt")
	handleError(err)
	length, err := io.WriteString(file, content)
	handleError(err)
	fmt.Println("Length of the file is", length)
	defer file.Close()
	readFile("./sample_file.txt")

}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	handleError(err)
	fmt.Println("Content of the file is: ")
	fmt.Println(string(databyte))
}
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
