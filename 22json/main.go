package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson :=
		`[
		{
			"first_name":"Clark",
			"last_name":"Kent",
			"hair_color":"black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name":"Wayne",
			"hair_color":"black",
			"has_dog": false
		}
	]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Fatal("Encountered error ", err)
	}
	log.Printf("Unmarshalled value %v", unmarshalled)

	//write json from struct
	var mySlice []Person
	flash := Person{
		FirstName: "Wally",
		LastName:  "West",
		HairColor: "Blonde",
		HasDog:    true,
	}
	wonder_woman := Person{
		FirstName: "Diana",
		LastName:  "Prince",
		HairColor: "Black",
		HasDog:    false,
	}
	mySlice = append(mySlice, flash, wonder_woman)
	newJson, err := json.MarshalIndent(mySlice, "", "    ")
	//newJson, err := json.Marshal(mySlice)
	if err != nil {
		log.Fatal("Encountered error ", err)
	}
	fmt.Println(string(newJson))
}
