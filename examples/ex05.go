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

    b3, _ := json.Marshal(d3)
    fmt.Println(string(b3))
	ioutil.WriteFile("ex05.json",b3,0644)

}
