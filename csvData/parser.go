package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Data struct {
	columns int
	data int //temp un-used -> pass int dataframe of converted csv
}

func GetCSVData(path string) int { //temp return int (cause currently returning number of columns)
	var d Data // d is of type struct Data

	f, err := os.Open(path) 
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f) //csv reader (pass in io reader)
	for { //try read from reader
		row, err := reader.Read() //reader returns row and error
		if err == io.EOF { //if end of file/row, stop reading
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		d.addData(len(row)) // access class method, add value to 'columns'
	}
	 return d.columns
}

func (d *Data) addData(c int) {
	d.columns = c // 'c' is the value passed in (len(row))
}

