package main

import (
	"github.com/Aidan-Zamfir/normalizer/cli"
)

// >>> go run . stand testdata.csv
// go run . stand testdata.csv -e ../../Desktop/test.csv

// >>> go run . nm testdata.csv
// go run . nm testdata.csv -e ../../Desktop/test.csv

func main() {
	cli.Execute()
}
