package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file"

	file, err := os.Create("./mylogfile.txt")
	checknilerr(err)
	length, err := io.WriteString(file, content)

	checknilerr(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./mylogfile.txt")
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	// if err != nil {
	// 	panic(err)
	// }
	checknilerr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))

}
func checknilerr(err error) {
	if err != nil {
		panic(err)
	}
}
