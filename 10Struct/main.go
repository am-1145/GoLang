package main

import "fmt"

func main() {
	fmt.Println("Structs in GoLang")

	// No Inheritence in golang; No super or parent
	Ankit := User{"Ankit", "am1145@g.com", true, 16}
	fmt.Println(Ankit)
	fmt.Printf("Ankit details are : %+v\n", Ankit)
	fmt.Printf("Ankit details are : %+v\n", Ankit.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
