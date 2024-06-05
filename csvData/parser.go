package csvData

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

//want something dynamic...

type UserData struct {
	First  []float64
	Second []float64
	Third  []float64
	Fourth []float64
}

func GetCSVData(path string) (*UserData, error) {

	c := &UserData{} //do like this usually

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err) //change err type
	}

	reader := csv.NewReader(f) //csv reader (pass in io reader)

	for i := 0; true; i++ {

		col, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err) //change
		}
		if i == 0 {
			continue
		}

		//CHNAGE FROM HERE
		first, err := strconv.ParseFloat(col[0], 64)
		if err != nil {
			return nil, err //write errors like this
		}

		second, err := strconv.ParseFloat(col[1], 64)
		if err != nil {
			return nil, err
		}

		third, err := strconv.ParseFloat(col[2], 64)
		if err != nil {
			return nil, err
		}

		fourth, err := strconv.ParseFloat(col[3], 64)
		if err != nil {
			return nil, err
		}

		c.First = append(c.First, first)
		c.Second = append(c.Second, second)
		c.Third = append(c.Third, third)
		c.Fourth = append(c.Fourth, fourth)

		//TO HERE

	}

	return c, nil //write this when have errors as return
}
