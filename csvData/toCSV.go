package csvData

import (
	"encoding/csv"
	//"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ToCSV(arr [][]float64) error {

	csvfile, err := os.Create("dataNM.csv")

	if err != nil {
		log.Fatal(err)
	}

	csvW := csv.NewWriter(csvfile)

	newArr := []string{}

	for _, i := range arr {
		for j := range i {
			x := i[j]
			s := strconv.FormatFloat(x, 'f', -1, 64)
			newArr = append(newArr, s)
			//_ = csvW.Write(s) //needs to be []string not 'string'
		}
		_ = csvW.Write(newArr)
		newArr = nil
	}

	csvW.Flush()
	err = csvfile.Close() //returns success (0) to err??
	if err != nil {
		log.Fatal(err)
	}
	return err
}

// WORKS:
//func ToSCV(arr [][]float64) {
//
//	for _, i := range arr {
//		for j := range i {
//			x := i[j]
//			s := strconv.FormatFloat(x, 'f', -1, 64)
//			log.Println(s)
//		}
//	}
//}
