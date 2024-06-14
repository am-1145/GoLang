package main

import "fmt"

func main() {
	fmt.Println("Welcome to class on Pointers")
	// var ptr *int
	// fmt.Println("Value of ptr is", ptr)

	myNumber := 23
	var ptr *int = &myNumber
	fmt.Println("Value of actual item is", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is", myNumber)
}
