package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"reflect"
	"time"
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

	//var mybyteary []byte

	d1a := Doc1{Path: "github.com/0studio/redisapi", Synopsis: "Purple", Stars: "0"}
	d1b := Doc1{Path: "github.com/aaasen/mycelium", Synopsis: "Blue", Stars: "0"}
	d1c := Doc1{Path: "github.com/stormasm/hola", Synopsis: "Green", Stars: "0"}

    docs1 := make([]Doc1, 0)
    docs1 = append(docs1,d1a)
    docs1 = append(docs1,d1b)
	docs1 = append(docs1,d1c)

	myjson := BuildJson(docs1)
	ioutil.WriteFile("ex07.json",myjson,0644)

	// now read the file back and get at the data
	var v interface{}
	data, err := ioutil.ReadFile("ex07.json")
	check(err)
	err = json.Unmarshal(data, &v)
	check(err)
	readJson(v)
}

func readJson(v interface{}) {

	fmt.Println("type:", reflect.TypeOf(v))

	m := v.([]interface{})

	for k, v := range m {

		fmt.Println(k,v)

		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println("process interface ary")
			//processAry(vv)
        case map[string]interface {}:
			fmt.Println("process map string interface ary")
            //processMap(vv)
		default:
			fmt.Print(k, " is of a type I don't know how to handle which is ")
            fmt.Println(reflect.TypeOf(v))
		}
	}
}

func BuildJson(repos []Doc1) ([]uint8){

	d3 := make(Doc2)
    d3["data"] = repos

	////////////////////////

	d4 := make(Doc2)
    d4["Repo"] = "garyburd/redigo"
    d4["Time"] = time.Now().String()

	docs := make([]Doc2,0)
	docs = append(docs,d4)

	d6 := make(Doc2)
	d6["meta"] = docs

	////////////////////////

	s := make([]interface{}, 2)
	s[0] = d3
	s[1] = d6

	//fmt.Println(s)

	b5, _ := json.Marshal(s)

	//fmt.Println(reflect.TypeOf(b5))

	return b5
	//fmt.Println(string(b5))
	//ioutil.WriteFile("ex07.json",b5,0644)

}
