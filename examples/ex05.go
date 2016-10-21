package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Doc1 struct {
	Path string
	Synopsis string
	Stars string
}

type Doc2 map[string]interface{}

func main() {

	d1 := Doc1{Path: "github.com/0studio/redisapi", Synopsis: "Purple", Stars: "0"}
	d2 := Doc1{Path: "github.com/aaasen/mycelium", Synopsis: "Blue", Stars: "10"}

    docs1 := make([]Doc1, 0)
    docs1 = append(docs1,d1)
    docs1 = append(docs1,d2)

	d3 := make(Doc2)
    d3["data"] = docs1

	////////////////////////

	d4 := make(Doc2)
    d4["Color"] = "Blue"
	d4["Name"] = "Iris"

    d5 := make(Doc2)
    d5["Color"] = "Purple"
	d5["Name"] = "Hb"

	docs := make([]Doc2,0)
	docs = append(docs,d4)
	docs = append(docs,d5)

	d6 := make(Doc2)
	d6["misc"] = docs

	////////////////////////

	s := make([]interface{}, 2)
	s[0] = d3
	s[1] = d6

	fmt.Println(s)

	b5, _ := json.Marshal(s)
	fmt.Println(string(b5))
	ioutil.WriteFile("ex05.json",b5,0644)

}
