package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tobgu/qframe"
)



func main() {

	var filepath string


	fmt.Print("Enter file path: ") //(testdata.csv)
    fmt.Scan(&filepath)

	cfile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	f := qframe.ReadCSV(cfile)
	fmt.Println(f)


	// csvData.GetCSVData(filepath)

	// data.Caller() //input Data struct 
	
}
