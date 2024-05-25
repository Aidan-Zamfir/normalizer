package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tobgu/qframe"
)

//parser temp not needed - remove later?

func main() {

	var filepath string


	fmt.Print("Enter file path: ") //(testdata.csv)
    fmt.Scan(&filepath)

	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	df := qframe.ReadCSV(f)

	fmt.Println(df)


	// data.Caller(df) //input Data struct 
	
}


// x, y := csvData.GetCSVData(filepath)
	// fmt.Println(x)
	// fmt.Println(" ")
	// fmt.Println(y)