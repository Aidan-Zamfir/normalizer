package csvData

import (
	//"encoding/csv"
	"log"
)

//func ToCSV(arr [][]float64) error {
//
//	csvfile, err := os.Create("dataNM.csv")
//
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	csvW := csv.NewWriter(csvfile)
//
//	for _, i := range arr {
//		log.Println(i)
//		s := strconv.FormatFloat(i, 'f', -1, 64)
//		_ = csvW.Write(s)
//	}
//	csvW.Flush()
//	err = csvfile.Close() //returns success (0) to err??
//	if err != nil {
//		log.Fatal(err)
//	}
//	return err
//}

func ToSCV(arr [][]float64) {

	for _, i := range arr {
		log.Println(i)
	}
}
