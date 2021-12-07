package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"
	fmt.Println("Length of array is: ", len(fruitList))
	for x := 0; x < len(fruitList); x++ {
		fmt.Printf("%v\t", fruitList[x])
	}
	var vegList [3]string = [3]string{"Tomato", "Potato", "Broccoli"}
	fmt.Println()
	for _, veggie := range vegList {
		fmt.Printf("%v\t", veggie)
	}
	fmt.Println()
	highScores := make([]int, 4)
	highScores[0] = 123
	highScores[1] = 987
	highScores[2] = 453
	highScores[3] = 231
	highScores = append(highScores, 456, 432, 972)
	fmt.Println(highScores)
	sort.Ints(highScores)
	fmt.Println(highScores)

	vegSlice := []string{}
	vegSlice = append(vegSlice, "Ginger", "Chilli", "Garlic")
	fmt.Println(vegSlice)

	courses := []string{"Python", "ReactJS", "NodeJS", "Swift", "Golang"}
	index := 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
