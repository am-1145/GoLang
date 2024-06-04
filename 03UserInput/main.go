package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating for our Pizza:\n ")

	// comma ok // err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
}
