package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(rw, "Hello World")
		//fmt.Fprintf(rw, "Hello World")
		if err != nil {
			fmt.Println("Encountered error: ", err)
		}
		fmt.Printf("Number of bytes written %d\n", n)
	})

	_ = http.ListenAndServe(":8080", nil)
	//http.ListenAndServe(":8080", nil)
}
