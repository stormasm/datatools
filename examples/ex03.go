package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Doc map[string]interface{}


func getDoc() Doc {

	name := []string{"michael","iris","hb","ken","stu","berk"}
	color := []string{"purple","blue","green","yellow","red","orange"}
	city := []string{"santafe","raton","taos","abq","bend","prineville"}
	month := []string{"jan","feb","mar","apr","may","june","july","aug","sep","october","november","december"}

	nameidx := rand.Int() % len(name)
	coloridx := rand.Int() % len(color)
	cityidx := rand.Int() % len(city)
	monthidx := rand.Int() % len(month)

	d := make(Doc)
    d["Name"] = name[nameidx]
	d["Color"] = color[coloridx]
	d["City"] = city[cityidx]
	d["Month"] = month[monthidx]

	return d
}


func main() {
	docs := make([]Doc, 0)
	for i := 0; i < 5; i++ {
		d := getDoc()
    	docs = append(docs,d)
	}
	b, _ := json.Marshal(docs)
    fmt.Println(string(b))
}
