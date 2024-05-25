package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)


var columns int


func GetCSVData(path string) int { //temp return int (cause currently returning number of columns)

	f, err := os.Open(path) 
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f) //csv reader (pass in io reader)
	// df := qframe.ReadCSV(f)

	for { //try read from reader
		col, err := reader.Read() //reader returns row and error
		if err == io.EOF { //if end of file/, stop reading
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//want to point to cloumns
		p := &columns
		*p = len(col)
		
		
	}
	
	return columns
}



