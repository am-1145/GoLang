package main

import "fmt"

func main() {
	fmt.Println("Welcome to method in golang")
	ankit := User{"ankit", "am", true, 19}
	ankit.GetStatus()
	ankit.Getmail()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}
func (u User) Getmail() {
	u.Email = "test@go.com"
	fmt.Println("Email of this user is: ", u.Email)
}
