package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func PrintInfo(a Animal) {
	fmt.Println("The Animal makes a noise of", a.Says(), "and has", a.NumberOfLegs(), "number of legs")
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Dog) Says() string {
	return "Woof"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}

func (g *Gorilla) Says() string {
	return "Ugh"
}
func main() {
	dog := Dog{
		Name:  "Gabba",
		Breed: "German Shepherd",
	}
	gorilla := Gorilla{
		Name:          "Mike",
		Color:         "Black",
		NumberOfTeeth: 42,
	}
	PrintInfo(&dog)
	PrintInfo(&gorilla)
}
