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

type Document struct {
	Key   string
	Value interface{}
}

func main() {

	d := Document{Key: "Name", Value: "Michael"}

	b, err := json.Marshal(d)
	check(err)
	fmt.Println(string(b))
}
