package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://chaicode.com/learn?course=golang"

func main() {
	fmt.Println("Welcome to url handling in GoLang")
	fmt.Println(myurl)

	//Parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are : %T\n", qparams)
	fmt.Println(qparams["course"])

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "chaicode.com",
		Path:     "/learn",
		RawQuery: "course=golang",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
