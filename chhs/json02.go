package main

import (
	// "bufio"
	// "bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processMap(vv map[string]interface {}) {
	fmt.Println("type = ", reflect.TypeOf(vv))
    // fmt.Println(vv)
	for k, u := range vv {
        fmt.Println(k)
        fmt.Println("\n\n\n")
        fmt.Println(u)
    }
}

func processAry(vv []interface{}) {
	fmt.Println("type of Big Array = ", reflect.TypeOf(vv))
	fmt.Println("length of Big Array = ", len(vv))
	for k, u := range vv {
		fmt.Println("type of Small Array = ", reflect.TypeOf(u))
		fmt.Println(k, u)

		switch vv := u.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			processAry(vv)
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

func processJson(v interface{}) {

	fmt.Println("type:", reflect.TypeOf(v))

	m := v.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			//processAry(vv)
        case map[string]interface {}:
            processMap(vv)
		default:
			fmt.Print(k, " is of a type I don't know how to handle which is ")
            fmt.Println(reflect.TypeOf(v))
		}
	}
}

func main() {
	var v interface{}
	data, err := ioutil.ReadFile("chem03.json")
	check(err)
	err = json.Unmarshal(data, &v)
	check(err)
	processJson(v)
}
