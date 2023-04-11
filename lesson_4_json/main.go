package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"usename"`
}

var DATA = `
{
	"id": 55,
	"price": 3000,
	"items": [
		{
			"name": "snowbord",
			"number": 1
		},
		{
			"name": "ball",
			"number": 4
		}
	]
}
`

type Order struct {
	Id    int    `json:"id"`
	Price int    `json:"price"`
	Items []Item `json:"items"`
	Like  string
}

type Item struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func marsh() {
	user1 := User{Name: "Ivan", Id: 2}
	bytes, _ := json.MarshalIndent(user1, "", "    ")
	fmt.Println(string(bytes))
}

func defaultMain() {
	var order Order
	err := json.Unmarshal([]byte(DATA), &order)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", order)
}

func main() {
	ReturningJSON()
}
