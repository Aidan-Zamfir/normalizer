package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

var columns int //stores columns (to be used when normalizing)

func GetCSVData(path string) int{ //temp return int (cause returning columns)

	f, err := os.Open("testdata.csv")
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f) //csv reader (pass in io reader)
	for { //try read from reader
		row, err := reader.Read() //reader returns row and error
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		columns = len(row)
	}
		return columns
}
