package main

import (
	"encoding/json"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Doc1 struct {
	Name   string
	Color interface{}
}

type Doc2 map[string]interface{}

func main() {

	d1 := Doc1{Name: "Michael", Color: "Purple"}

	b1, err := json.Marshal(d1)
	check(err)
	fmt.Println(string(b1))

	d2 := make(Doc2)
    d2["Color"] = "Blue"
	d2["Name"] = "Iris"
	b2, _ := json.Marshal(d2)
	fmt.Println(string(b2))
}
