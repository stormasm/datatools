package main

import (
	"fmt"
	"github.com/go-resty/resty"
)

func main() {
	resty.SetOutputDirectory("/tmp57")

	_, err := resty.R().
		SetOutput("medreview.json").
		Get("https://chhs.data.ca.gov/api/views/pch7-48qc/rows.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("done")
}
