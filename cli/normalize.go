package cli

import (
	"log"

	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/spf13/cobra"
)

var normalizeCmd = &cobra.Command{
	Use:   "nm",
	Short: "Will return normalized values as X file form", //decide
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		d, err := csvData.GetCSVData(args[0])
		if err != nil {
			log.Fatal(err)
		}

		//temp fix for testing:
		collist := [][]float64{d.First, d.Second, d.Third, d.Fourth}
		//Currently only passing in one slice of flt64 (1 column) -> not entire dataset
		//call function within parser that iterates over columns and returns each for equation use

		for i := range collist {
			result := data.MinMax(collist[i])
			log.Println(result)
			//add function that saves them somewhere else as CSV
		}
		//result := data.MinMax(d.First)
		//log.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->")
}
