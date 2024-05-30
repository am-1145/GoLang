package main

import "fmt"

const LoginToken string = "kbkasfas"

func main() {
	var username string = "Ankit"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)
	var check bool = false
	fmt.Println(check)
	fmt.Printf("Variable is of type: %T \n", check)

	var another string
	fmt.Println(another)
	fmt.Printf("Variable is of type: %T \n", another)

	// no var style
	numberofUser := 300000.0
	fmt.Println(numberofUser)
	fmt.Printf("Variable is of type: %T \n", numberofUser)

}

// If we declare any variable supposxe LoginToken then in go the first
// letter of the variable should be capital letter if we want to use it
// as public otherwise it will be private to that package
