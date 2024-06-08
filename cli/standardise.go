package cli

import (
	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/spf13/cobra"
	"log"
)

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
			go data.MinMax(da[i], c)
			result := <-c
			cols = append(cols, result)
		}

		err = csvData.ToCSV(head, cols, 1)
		if err != nil {
			log.Fatal(err)
		}

		//res := "dataST.csv"
		//ExportToFile(res, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
	//standardiseCmd.PersistentFlags().String("filepath", "export", "Export to filepath -> (provide path)")
}
