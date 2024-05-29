package cli

import (
	"fmt"
	"log"

	"github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/Aidan-Zamfir/normalizer/data"
	"github.com/spf13/cobra"
)

var standardiseCmd = &cobra.Command{
	Use: "stand", 
	Short: "Will return standardised values as X file form", //decide
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		d, err := csvData.GetCSVData(args[0])
		if err != nil {
				log.Fatal(err)
			}

		fmt.Println(" ")

		result := data.Standardise(d.First)
		fmt.Println(result, "<- data std")
		// ExportToFile(result, exportfilepath)
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->") 
}