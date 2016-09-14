package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

var columnames []string

func processData(i int,row []string) {
    if i == 0 {
		columnames = make([]string, len(row))
		copy(columnames, row)
    } else {
        fmt.Println("Data")
        buildJson(columnames,row)
    }
}

func buildJson(column []string, row []string) {
	for i, colname := range column {
		fmt.Println(i,colname)
	}
	for i, record := range row {
		fmt.Println(i,record)
	}
}

func main() {

	csvfile, err := os.Open("chem02.csv")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i, row := range rawCSVdata {
		fmt.Println()
		fmt.Println(i,row)
		fmt.Println()
		processData(i,row)
	}
}
