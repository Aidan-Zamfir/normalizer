package csvData

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// parse csv,
// find len (number of columns)
// itterate through itrms in column and append to slice?

type userData struct {
	columns int
	data []float64
}

func GetCSVData(path string) userData { //temp return int (cause currently returning number of columns)
	// var columns int = 0
	var c userData

	f, err := os.Open(path) 
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f) //csv reader (pass in io reader)


	for { 
		col, err := reader.Read() 
		if err == io.EOF { 
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		c.columns = len(col)
	}
	fmt.Println(c.columns)
	
	return userData{}
}