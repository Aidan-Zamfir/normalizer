package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Titles struct {
	headers []string
}

func GetCSVData(path string) (table [][]float64, headers [][]string, err error) {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f) //csv reader (pass in io reader)
	table = [][]float64{}
	headers = [][]string{}

	for i := 0; ; i++ {

		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}
		if i == 0 {
			headers = make([][]string, len(row))
			table = make([][]float64, len(row))
			for j, row := range row {
				headers[j] = append(headers[j], row)
			}
			continue
		}
		for j, row := range row {
			cell, err := strconv.ParseFloat(row, 64)
			if err != nil {
				return nil, nil, err
			}
			table[j] = append(table[j], cell)
		}
	}
	//log.Println(headers)
	return table, headers, err //write this when have errors as return
}
