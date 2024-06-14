package main

import "fmt"

func main() {
	fmt.Println("If else in golannng")
	loginCount := 1
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "watch out"
	} else {
		result = "exactly 10 "
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}
