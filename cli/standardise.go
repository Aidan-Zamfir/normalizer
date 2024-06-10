package cli

import (
	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/spf13/cobra"
	"log"
)

var exportFilePath string

var standardiseCmd = &cobra.Command{
	Use:   "stand",
	Short: "Will return standardised values as .csv file", //decide
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan []float64)

		da, head, err := csvData.GetCSVData(args[0])
		if err != nil {
			log.Fatal(err)
		}

		cols := [][]float64{}

		for i := range da {
			go data.Standardise(da[i], c)
			result := <-c
			cols = append(cols, result)
		}
		if exportFilePath != "" {
			err = csvData.ToCSV(head, cols, exportFilePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = csvData.ToCSV(head, cols, "./DataST.csv")
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
	standardiseCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to filepath -> (provide path) | if empty -> ./DataST.csv")
}

// go run . stand testdata.csv -e ../../Desktop/test.csv
