package main

import (
	// "bufio"
	// "bytes"
	"io/ioutil"
	"fmt"
    "encoding/json"
    "reflect"
)

func check(e error) {
	if e != nil {
		panic(e)
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
            processAry(vv)
        default:
            fmt.Println(k, "is of a type I don't know how to handle")
        }
    }

/*
    switch v := i.(type) {
    case int:
        fmt.Println("twice i is", v*2)
    case float64:
        fmt.Println("the reciprocal of i is", 1/v)
    case string:
        h := len(v) / 2
        fmt.Println("i swapped by halves is", v[h:]+v[:h])
    default:
        fmt.Println("i isn't one of the types above")
    }
*/
}

func main() {
    var v interface{}
	data, err := ioutil.ReadFile("chem03.json")
	check(err)
    err = json.Unmarshal(data, &v)
    check(err)
    processJson(v)
}
