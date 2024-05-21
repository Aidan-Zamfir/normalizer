package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	var filepath string
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})


	fmt.Print("Enter file path: ") //(testdata.csv)
    fmt.Scan(&filepath)

	x := csvData.GetCSVData(filepath)
	fmt.Println(x, "<-- columns")

	data.Caller() //input Data struct
	
	
}
