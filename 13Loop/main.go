package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop")
	days := []string{"Sun", "mon", "tue", "wed", "thu", "fri", "sat"}
	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, days := range days {
	// 	fmt.Printf("Index is %v and value is %v\n", index, days)
	// }
	// for _, days := range days {
	// 	fmt.Printf("Index is  and value is %v\n", days)
	// }

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 2 {
			goto lco

		}
		fmt.Println(rogueValue)
		rogueValue++
	}

lco:
	fmt.Println("Jumping at 1")

}
