package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func processData(i int,row []string) {
    if i == 0 {
        readColumns(row)
    } else {
        fmt.Println("Data")
        readData(row)
    }
}

func readColumns(row []string) {
    fmt.Println(row)
}

func readData(row []string) {
    fmt.Println(row)
}

func main() {

	csvfile, err := os.Open("chem02.csv")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)

	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for i, row := range rawCSVdata {
		//fmt.Println(i,row)
        //fmt.Println("\n")
        processData(i,row)
	}
}
