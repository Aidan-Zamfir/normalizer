package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func GetCSVData(path string) (table [][]float64, err error) {

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f) //csv reader (pass in io reader)
	table = [][]float64{}

	for i := 0; ; i++ {

		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if i == 0 {
			table = make([][]float64, len(row))
			continue
		}
		for j, row := range row {
			cell, err := strconv.ParseFloat(row, 64)
			if err != nil {
				return nil, err
			}
			table[j] = append(table[j], cell)
		}
	}
	return table, nil //write this when have errors as return
}
