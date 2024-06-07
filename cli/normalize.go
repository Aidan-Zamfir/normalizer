package cli

import (
	"github.com/Aidan-Zamfir/normalizer/data"
	"log"

	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/spf13/cobra"
)

var normalizeCmd = &cobra.Command{
	Use:   "nm",
	Short: "Will return normalized values as .csv file", //decide
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		da, head, err := csvData.GetCSVData(args[0])
		if err != nil {
			log.Fatal(err)
		}

		cols := [][]float64{}

		for i := range da {
			result := data.MinMax(da[i])
			cols = append(cols, result)
		}

		err = csvData.ToCSV(head, cols, 0)
		if err != nil {
			log.Fatal(err)
		}

		//res := "dataNM.csv"
		//ExportToFile(res, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)
	//normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to filepath -> (provide path)")
}
