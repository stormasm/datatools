package main

import (
	"fmt"
    "github.com/go-resty/resty"
)

func main() {

// Setting output directory path, If directory not exists then resty creates one!
// This is optional one, if you're planning using absoule path in
// `Request.SetOutput` and can used together.
resty.SetOutputDirectory("/tmp57")

// HTTP response gets saved into file, similar to curl -o flag
_, err := resty.R().
          SetOutput("medreview.json").
          Get("https://chhs.data.ca.gov/api/views/pch7-48qc/rows.json")

if err != nil {
    fmt.Println(err)
}

fmt.Println("done")
}
