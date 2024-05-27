package cli

import (
	// "github.com/Aidan-Zamfir/normalizer/csvData"
	"github.com/spf13/cobra"
)

var standardiseCmd = &cobra.Command{
	Use: "stand", 
	Short: "Will return standardised values as X file form", //decide
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// - arg MUST be file path to csv file.
		// - this func calls a parser method to convert file passed in
		// - based on user, will nm or stand data
		// - return new data file



		//Call parser method first-> then pass into Stnd func
		// data := csvData.GetCSVData(args) ?? or qFrame?
		
		// result := data.Standardise(args)
		// fmt.Println(result)
		// ExportToFile(result, exportfilepath)
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->") 
}