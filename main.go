package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aidan-Zamfir/normalizer/csvData"
)



func main() {

	var filepath string


	fmt.Print("Enter file path: ") //(testdata.csv)
    fmt.Scan(&filepath)

	_, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	// f := qframe.ReadCSV(cfile)
	// fmt.Println(f)


	x := csvData.GetCSVData(filepath)
	fmt.Println(x)

	// data.Caller() //input Data struct 
	
}
