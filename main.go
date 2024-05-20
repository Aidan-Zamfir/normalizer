package main

import (
	"fmt"

	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
)

func main() {

	var filepath string

	fmt.Print("Enter file path: ") //(testdata.csv)
    fmt.Scan(&filepath)

	x := csvData.GetCSVData(filepath)
	fmt.Println(x, "<-- columns")

	data.Caller()
	
	
}

