package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/am-1145/mongoapi/router"
)

func main() {
	fmt.Println("MONGO API")
	r := router.Router()
	fmt.Println("Server is Getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000")
}
