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
		fmt.Println(args[0])
		d, err := csvData.GetCSVData(args[0])
		if err != nil {
				log.Fatal(err)
			}

		result := data.Standardise(d.First)
		fmt.Println(result)
		// ExportToFile(result, exportfilepath)
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->") 
}