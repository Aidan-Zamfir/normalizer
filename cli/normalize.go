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

		cols := [][]float64{}

		//temp fix for testing:
		collist := [][]float64{d.First, d.Second, d.Third, d.Fourth}
		//call function within parser that iterates over columns and returns each for equation use

		for i := range collist {
			result := data.MinMax(collist[i])
			cols = append(cols, result)
			//log.Println(cols)
			//add function that saves them somewhere else as CSV
		}

		err = csvData.ToCSV(cols)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->")
}
