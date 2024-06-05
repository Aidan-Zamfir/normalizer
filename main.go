package main

import (
	//"github.com/Aidan-Zamfir/normalizer/cli"
	"encoding/csv"
	"io"
	"log"
	"os"
	//"strconv"
)

func main() {
	//cli.Execute()

	//TEST FUNC:
	err := getD("testdata.csv")
	if err != nil {
		log.Fatal(err)
	}
}

func getD(path string) error {

	//colData := [][]float64{} //holds array of arrays flt64
	//var rowLen int = 100 //temp value 100, reset after loop starts

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := csv.NewReader(f) //csv reader (pass in io reader)

	//for i := 0; true; i++ { // need length
	for {
		var rowLen int = 100

		count := 0
		x := 0 //first iter is over column 0
		log.Println(x, "x start")
		for i := range rowLen { // need length
			log.Println(rowLen, "rowlen first")

			row, err := reader.Read()
			rowLen = len(row)

			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}
			if i == 0 {
				continue
			}

			log.Println(row[x], "ELEMENT")

			//j, err := strconv.ParseFloat(row[x], 64) // j is the float64 value of the element at row[x]
			//if err != nil {
			//	return err
			//}

			//colData = append(colData, []float64{j})
			//log.Println(colData)
			log.Println(rowLen, "rowlen end of loop")
			log.Println(x, "x end of loop")
		}
		x++ // x adds for items in columns until reach no more items
		count++
		log.Println(rowLen, "rowLen LAST")
		log.Println(count, "count end")
		log.Println(x, "x end")
		if count > rowLen {
			break
		}
	}
	return err
}

// >>> go run . stand testdata.csv

//at the point where im iterating over one column fully, but not moving on to the next
