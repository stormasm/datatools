package main

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"reflect"
	"time"
	//"github.com/stormasm/reflections"
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
	ioutil.WriteFile("ex08.json",myjson,0644)

	// now read the file back and get at the data
	var v interface{}
	data, err := ioutil.ReadFile("ex08.json")
	check(err)
	err = json.Unmarshal(data, &v)
	check(err)
	readJson(v)
}

func readJson(v interface{}) {

	//fmt.Println("type:", reflect.TypeOf(v))

	m := v.([]interface{})

	for k, v := range m {

		//fmt.Println(k)

		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			fmt.Println("process interface ary")
			//processAry(vv)
        case map[string]interface {}:
			//fmt.Println("process map string interface ary")
            processMap(vv)
		default:
			fmt.Print(k, " is of a type I don't know how to handle which is ")
            fmt.Println(reflect.TypeOf(v))
		}
	}
}

func processDataMap(t interface{}) {
	fmt.Println(t)
	//mytype := reflect.TypeOf(t).Kind()
	value := reflect.ValueOf(t)
	//fmt.Println(mytype, value)
	fmt.Printf("%+v\n", value)

	fmt.Println(reflect.Indirect(value))
/*
        for i := 0; i < s.Len(); i++ {
            mymap := s.Index(i)
			fmt.Println(mymap)
        }
*/

/*
	mystruct := reflect.TypeOf(t).Kind()
	fmt.Println(mystruct)
	value := reflect.ValueOf(t)
	fmt.Println(value.String())
*/

/*
	fmt.Println("map = ",t)
	mydoc1 := reflect.TypeOf(t).Kind()
	fmt.Println(mydoc1)
	value := reflect.ValueOf(t).Interface()
	fmt.Println(value)
	fields, _ := reflections.Fields(value)
	fmt.Println(fields)
	valuex, _ := reflections.Items(mydoc1)
	fmt.Println(valuex)
*/
}


func processDataKey(t interface{}) {
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(t)

        for i := 0; i < s.Len(); i++ {
            mymap := s.Index(i)
			processDataMap(mymap)
        }
    }
}

func processMetaKey(t interface{}) {
    switch reflect.TypeOf(t).Kind() {
    case reflect.Slice:
        s := reflect.ValueOf(t)

        for i := 0; i < s.Len(); i++ {
			mymap := s.Index(i)
            fmt.Println(mymap)
        }
    }
}

func processMap(vv map[string]interface {}) {
	//fmt.Println("map type = ", reflect.TypeOf(vv))
    // fmt.Println(vv)
	for key, value := range vv {
        //fmt.Println("key = ",key)
		//fmt.Println("value = ",value)
		switch key {
		case "data":
			processDataKey(value)
		case "meta":
			processMetaKey(value)
		}
        //fmt.Println("type of inner map = ", reflect.TypeOf(u))
        //fmt.Println("\n\n\n")
        //processInnerMap(u)
    }
}

func processInnerMap(v interface {}) {
	fmt.Println("map 222 type = ", reflect.TypeOf(v))
/*
    m := v.(map[string]interface{})

	for k, v := range m {
        fmt.Println("Inner key = ", k, v)
	}
*/
}

func processColumns(vv interface {}){

	fmt.Println("column type:", reflect.TypeOf(vv))
    //fmt.Println("value type:", reflect.ValueOf(vv))

    v := reflect.ValueOf(vv)

    //fmt.Println("type:", v.Type())
    fmt.Println("column kind", v.Kind())

    for i := 0; i < v.Len(); i++ {
        mymap := v.Index(i)
        fmt.Println(mymap)
        fmt.Println("\n\n")


     }


//    fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
//    fmt.Println("value:", v.Float())




    //fmt.Println(m)

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
