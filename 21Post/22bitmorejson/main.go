package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourse := []course{
		{"ReactJS", 299, "LCO", "a12345", []string{"web-dev", "js"}},
		{"MERN", 199, "LCO", "b12345", []string{"fullstack", "js"}},
		{"Angular", 299, "LCO", "c12345", nil},

		// package this data as json data

	}
	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	  {
        "coursename": "ReactJS",
        "Price": 299,
        "Platform": "LCO",
        "Password": "a12345",
        "Tags": ["web-dev","js"]
        }
	`)
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON Was Validated")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON Was Not Validated")
	}
	// Some cases where you just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v , value is %v and Type is:  %T\n", k, v, v)
	}

}
