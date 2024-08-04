package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
}

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	IsStudent bool     `json:"isStudent"`
	Courses   []string `json:"courses"`
	Address   Address  `json:"address"`
}

func main() {
	jsonString := `{"name":"John","age":30,"isStudent":false,"courses":["Math","Science","History"],"address":{"street":"123 Main St","city":"Anytown"}}`

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", person)

	personObject := Person{
		Name:      "John",
		Age:       30,
		IsStudent: false,
		Courses:   []string{"Math", "Science", "History"},
		Address: Address{
			Street: "123 Main St",
			City:   "Anytown",
		},
	}
	jsonData, err := json.Marshal(personObject)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)
}
