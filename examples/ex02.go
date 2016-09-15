package main

import (
	// "bufio"
	// "bytes"
	"encoding/json"
	"fmt"
	//"os"
	//"io/ioutil"
	//"reflect"
	//"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Document struct {
	key   string
	//value interface{}
    value string
}

func main() {

	d := Document{"Name", "Michael"}

	b, err := json.Marshal(d)
	check(err)
	//os.Stdout.Write(b)
	fmt.Println(string(b))
}
