package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Data struct {
	columns int
	data int //temp un-used
}

// var columns int //stores columns (to be used when normalizing)

func GetCSVData(path string) int { //temp return int (cause returning columns)
	var d Data

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

		d.addData(len(row))
	}
	 return d.columns
}

func (d *Data) addData(c int) {
	d.columns = c
}


<<<<<<< HEAD
=======

>>>>>>> dev
