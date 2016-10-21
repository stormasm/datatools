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

	b1, err := json.Marshal(d1)
	check(err)
	fmt.Println(string(b1))

	d3 := make(Doc2)
    d3["Color"] = "Blue"
	d3["Name"] = "Iris"
	b2, _ := json.Marshal(d3)
	fmt.Println(string(b2))

    // create a slice of doc2

    d4 := make(Doc2)
    d4["Color"] = "Purple"
	d4["Name"] = "Hb"


    docs1 := make([]Doc1, 0)
    docs1 = append(docs1,d1)
    docs1 = append(docs1,d2)

	docs := make([]Doc2,0)
	docs = append(docs,d3)
	docs = append(docs,d4)

    b3, _ := json.Marshal(docs)
    fmt.Println(string(b3))
	ioutil.WriteFile("ex04a.json",b3,0644)

	b4, _ := json.Marshal(docs1)
    fmt.Println(string(b4))
	ioutil.WriteFile("ex04b.json",b4,0644)

	s := make([]interface{}, 2)
	s[0] = docs1
	s[1] = docs

	fmt.Println(s)

	b5, _ := json.Marshal(s)
	fmt.Println(string(b5))
	ioutil.WriteFile("ex04c.json",b5,0644)

}
