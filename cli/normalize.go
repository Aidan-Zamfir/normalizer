package cli

import (
	"github.com/spf13/cobra"
)

var normalizeCmd = &cobra.Command{
	Use: "nm", 
	Short: "Will return normalized values as X file form", //decide
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//result := Normalize() (function)
		//fmt.Println(result)
		//ExportToFile(result, exportfilepath)
	},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)
	//need to add file path functionality
	// normalizeCmd.Flags().StringVarP(&exportFilePath, "export", "e", "Export to file ->") 
}