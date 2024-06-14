package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Map")

	languages := make(map[string]string)
	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages are", languages)
	fmt.Println("JS full form", languages["JS"])
	delete(languages, "RB")
	fmt.Println("List of all languages are", languages)

	// Loops in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
