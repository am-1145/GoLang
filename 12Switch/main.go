package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switchcase in golang")

	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got one")
	case 2:
		fmt.Println("You got two")
	case 3:
		fmt.Println("You got three")
		fallthrough
	case 4:
		fmt.Println("You got four")
	case 5:
		fmt.Println("You got five")
	case 6:
		fmt.Println("You got six")
	default:
		fmt.Println("What was that!")
	}
}
