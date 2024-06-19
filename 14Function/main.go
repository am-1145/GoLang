package main

import "fmt"

func greet() {
	fmt.Println("hi")
}
func main() {
	greet()
	fmt.Println("Welcome to function")
	result := adder(3, 5)
	fmt.Println("The result is", result)
	res := proadd(1, 2, 3, 4)
	fmt.Println("The result of proadd is", res)
}
func adder(val1 int, val2 int) int {
	return val1 + val2
}
func proadd(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}
