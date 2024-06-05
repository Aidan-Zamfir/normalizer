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

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	reader := csv.NewReader(f) //csv reader (pass in io reader)

	//colData := [][]float64{} //holds array of arrays flt64
	//var rowLen int = 100 //temp value 100, reset after loop starts

	//for i := 0; true; i++ { // need length

	for {

		var (
			rowLen = 100 //temp value 100 -> is set to correct length of row AFTER FIRST loop
			count  = 0   //increments with each loop cycle until > rowLen
			x      = 0   //first iter is over column 0 (increments with each loop cycle)
			y      int
		)

		for i := range rowLen {
			log.Println("itteration: ", i)

			row, err := reader.Read()

			rowLen = len(row)

			//ISSUE HERE - READER NOT RE-LOADING AND RE-READING | OR LOOPING ISSUES
			if err == io.EOF {
				log.Println("BREAK EOF")
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
		}
		x++ // x adds for items in columns until reach no more items
		count++

		if count > y {
			break
		}
	}
	return err
}

// >>> go run . stand testdata.csv

//at the point where im iterating over one column fully, but not moving on to the next
