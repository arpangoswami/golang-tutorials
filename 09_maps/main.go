package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps")
	lang := make(map[string]string)
	lang["JS"] = "Javascript"
	lang["RB"] = "Ruby"
	lang["PY"] = "Python"
	fmt.Println(lang)
	fmt.Println("JS stands for:", lang["JS"])
	delete(lang, "RB")
	fmt.Println(lang)

	for key, value := range lang {
		fmt.Printf("For key %v value %v\n", key, value)
	}
}
