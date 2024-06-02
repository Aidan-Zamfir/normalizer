package cli

import (
	"log"

	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/spf13/cobra"
)

var standardiseCmd = &cobra.Command{
	Use:   "stand",
	Short: "Will return standardised values as X file form", //decide
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		d, err := csvData.GetCSVData(args[0])
		if err != nil {
			log.Fatal(err)
		}

		collist := [][]float64{d.First, d.Second, d.Third, d.Fourth}

		for i := range collist {
			result := data.Standardise(collist[i])
			log.Println(result)
			//add function that saves them somewhere else as CSV
		}
		//result := data.Standardise(d.First)
		//log.Println(result, "<- data std")
		// ExportToFile(result, exportfilepath)
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->")
}
