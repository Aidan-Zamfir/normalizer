package main

import (
	"github.com/Aidan-Zamfir/normalizer/cli"
)

func main() {
	cli.Execute()

}

// >>> go run . stand testdata.csv

// main -> cli/root.go -> execute()  				WORKS
// execute() -> init() 								WORKS
// init() -> normalizeCmd() | standardiseCmd() 		WORKS

// NM:
// -> csvData/parser.go -> getCSVData(FILEPATH)		WORKS
// •parrses csv file -> returns array struct (d)	WORKS
// -> data/equations -> minmax(d)				  	  /
// minmac(d) -> minimum(d)						 	  /
//			 -> maximum(d)			  				  /
// •normalize data -> retuns []float64{}			  /
// NEEDS TO RETURN DATA AND CONVERT TO CSV			 (?)

// STAND:
// -> csvData/parser.go -> getCSVData(FILEPATH)		WORKS
// •parrses csv file -> returns array struct (d)	WORKS
// -> data/equations -> Standardise(d)				  /
// Standardise(d) -> mean(d)						  /
//				  -> standardDeviation(d)			  /
// •standardises data -> retuns []float64{}			  /
// NEEDS TO RETURN DATA AND CONVERT TO CSV			 (?)
