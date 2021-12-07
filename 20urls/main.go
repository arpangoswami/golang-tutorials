package main

import (
	"fmt"
	"log"
	"net/url"
)

const imag_url string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hjg234j2g3"

func main() {
	fmt.Println("Welcome to extracting params based on urls")
	result, err := url.Parse(imag_url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of qparams is %T\n", qparams)
	for _, val := range qparams {
		fmt.Println(val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=arpan",
	}

	urlString := partsOfUrl.String()
	fmt.Println(urlString)
}
