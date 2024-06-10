package cli

import (
	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/spf13/cobra"
	"log"
)

var normalizeCmd = &cobra.Command{
	Use:   "nm",
	Short: "Will return normalized values as .csv file", //decide
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c := make(chan []float64)

		da, head, err := csvData.GetCSVData(args[0])
		if err != nil {
			log.Fatal(err)
		}

		cols := [][]float64{}

		for i := range da {
			go data.MinMax(da[i], c)
			result := <-c
			//result := data.MinMax(da[i])
			cols = append(cols, result)
		}

		if exportFilePath != "" {
			err = csvData.ToCSV(head, cols, exportFilePath)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			err = csvData.ToCSV(head, cols, "./DataNM.csv")
			if err != nil {
				log.Fatal(err)
			}
		}

	},
}

func init() {

	rootCmd.AddCommand(normalizeCmd)
	normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to filepath -> (provide path) | if empty -> ./DataNM.csv")
}
