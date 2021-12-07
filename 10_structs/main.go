package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to structs")
	arpan := User{Name: "Arpan", Email: "arpan@gmail.com", Status: true, Age: 21}
	fmt.Println(arpan)
	fmt.Printf("Arpan details are %+v\n", arpan)
	fmt.Printf("%v's email is %v\n", arpan.Name, arpan.Email)
}
