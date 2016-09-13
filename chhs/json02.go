package main

import (
	// "bufio"
	// "bytes"
	"io/ioutil"
	// "fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	_, err := ioutil.ReadFile("chem01.json")
	check(err)
}
