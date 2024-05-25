package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/tobgu/qframe"
)





func GetCSVData(path string) (int, qframe.QFrame){ //temp return int (cause currently returning number of columns)
	var columns int = 0

	f, err := os.Open(path) 
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(f) //csv reader (pass in io reader)
	df := qframe.ReadCSV(f)
	
	for { 
		col, err := reader.Read() 
		if err == io.EOF { 
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// p := &columns
		// *p = len(col)
		columns = len(col)
	}
	
	return columns, df
}



