package main

import (
	"github.com/Aidan-Zamfir/normalizer/csvData"
)


func main() {
	// cli.Execute()
	csvData.GetCSVData("testdata.csv")
}
	

	// var filepath string


	// fmt.Print("Enter file path: ") //(testdata.csv)
    // fmt.Scan(&filepath)

	// f, err := os.Open(filepath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// df := qframe.ReadCSV(f)

	// fmt.Println(df)


	// data.Caller(df) //input Data struct 
	



// x, y := csvData.GetCSVData(filepath)
	// fmt.Println(x)
	// fmt.Println(" ")
	// fmt.Println(y)