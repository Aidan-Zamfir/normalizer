package main

import (
	"github.com/Aidan-Zamfir/normalizer/cli"
)

// >>> go run . stand test/testdata.csv
// go run . stand test/testdata.csv -e ../../Desktop/test.csv

// >>> go run . nm test/testdata.csv
// norm nm test/testdata.csv -e ../../Desktop/test.csv

func main() {
	cli.Execute()
}
