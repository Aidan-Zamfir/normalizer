package cli

import (
	"github.com/spf13/cobra"
)

var standardiseCmd = &cobra.Command{
	Use: "stand", 
	Short: "Will return standardised values as X file form", //decide
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	rootCmd.AddCommand(standardiseCmd)
}