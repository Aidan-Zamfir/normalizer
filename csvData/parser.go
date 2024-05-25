package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Data struct {
	columns int
	data float64//temp un-used -> pass int dataframe of converted csv
}

func GetCSVData(path string) int { //temp return int (cause currently returning number of columns)
	var d Data // d is of type struct Data

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

		d.addCol(len(col)) // access class method, add value to 'columns'
		
	}
	// d.addData(df) issue here -> need to create type of DF so it can go in struct (or just dont use struct)
	 return d.columns
}

func (d *Data) addCol(c int) {
	d.columns = c // 'c' is the value passed in (len(row))
}
func (d *Data) addData(df float64) {
	d.data = df // 'c' is the value passed in (len(row))
}

