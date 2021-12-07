package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Welcome to web request")
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type of resposne is %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Response received is ", string(databytes))
}
