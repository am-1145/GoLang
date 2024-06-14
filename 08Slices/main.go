package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitList = []string{"apple", "mango"}
	fmt.Printf("Type of fruitless is %T\n", fruitList)

	fruitList = append(fruitList, "banana")
	fmt.Println("fruitlist", fruitList)

	fruitList = append(fruitList[:2])
	fmt.Println("fruitlist", fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 944
	highScores[2] = 254
	highScores[3] = 264
	// highScores[4] = 465

	highScores = append(highScores, 555, 666)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// Remove value baseed on index

	var courses = []string{"React", "js", "swift", "go", "python"}

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
