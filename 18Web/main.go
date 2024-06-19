package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://chaicode.com/"

func main() {
	fmt.Println("Welcome to web request in GoLang")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close() // Caller's responsibilty to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(databytes))
}
