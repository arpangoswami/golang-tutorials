package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the Home page")
}

//About is the about route handler
func About(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the About page: %d", addValues(2, 2))
}

//add values
func addValues(x, y int) int {
	return (x + y)
}
func dividevalues(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("cannot divide by zero")
	}
	return x / y, nil
}
func Divide(rw http.ResponseWriter, r *http.Request) {
	res, err := dividevalues(100.0, 10.0)
	if err != nil {
		fmt.Fprintf(rw, "Cannot divide by zero error %v", err.Error())
		return
	}
	fmt.Fprintf(rw, "This is the divide page %v", res)
}
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)
	fmt.Println(fmt.Printf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
