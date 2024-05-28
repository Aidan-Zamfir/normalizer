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
// -> csvData/parser.go -> getCSVData(FILEPATH)		WORKS
// parrses csv file -> returns struct (d)			WORKS
// -> data/equations -> standardise(d)				  /

