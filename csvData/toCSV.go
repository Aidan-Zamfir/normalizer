package csvData

import (
	"encoding/csv"
	"os"
	"strconv"
)

// TO DO:
// Re-insert column names (get from parser?)

// t int,
//func ToCSV(head [][]string, arr [][]float64, filePath string) error {
//	//var filepath string
//
//	if filePath != "" {
//		csvFilePath, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//		if err != nil {
//			log.Println("Error opening file:", err)
//			return err
//		}
//		defer csvFilePath.Close()
//
//		//csvFilePath, err := os.Create(filePath) //new csv file saved to local device
//
//		//if err != nil {
//		//	return err
//		//}
//
//		csvW := csv.NewWriter(csvFilePath) // writer object csvFilePath -> where to write to
//		newArr := []string{}               // temp storage for string
//
//		for _, i := range head {
//			for j := range i {
//				x := i[j]
//				newArr = append(newArr, x) //need to do newline
//
//			}
//		}
//		_ = csvW.Write(newArr)
//		newArr = nil
//
//		for _, i := range arr {
//			for j := range i {
//				x := i[j]
//				s := strconv.FormatFloat(x, 'f', -1, 64)
//				newArr = append(newArr, s)
//			}
//			_ = csvW.Write(newArr)
//			newArr = nil //reset array after each loop (after being written to csv)
//		}
//
//		csvW.Flush()
//
//		//defer csvFilePath.Close()
//
//	}
//	return nil
//}

func ToCSV(head [][]string, arr [][]float64, filepath string) error {

	csvFile, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) //new csv file saved to local device
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
	_ = csvW.Write(newArr)
	newArr = nil

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
