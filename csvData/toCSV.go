package csvData

import (
	"encoding/csv"
	"os"
	"strconv"
)

// TO DO:
// Re-insert column names (get from parser?)

func ToCSV(head [][]string, arr [][]float64, t int) error {
	var filepath string

	if t == 0 {
		filepath = "dataNM.csv"
	} else if t == 1 {
		filepath = "dataST.csv"
	}

	csvFile, err := os.Create(filepath) //new csv file saved to local device

	if err != nil {
		return err
	}

	csvW := csv.NewWriter(csvFile) // writer object
	newArr := []string{}           // temp storage for string

	for _, i := range head {
		for j := range i {
			x := i[j]
			newArr = append(newArr, x) //need to do newline
		}
	}

	for _, i := range arr {
		for j := range i {
			x := i[j]
			s := strconv.FormatFloat(x, 'f', -1, 64)
			newArr = append(newArr, s)
		}
		_ = csvW.Write(newArr)
		newArr = nil //reset array after each loop (after being written to csv)
	}

	csvW.Flush()
	err = csvFile.Close() //returns success (0) to err
	if err != nil {
		return err
	}
	return err
}
