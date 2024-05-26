package cli

import (
	"github.com/spf13/cobra"
)

var normalizeCmd = &cobra.Command{
	Use: "nm", 
	Short: "Will return normalized values as X file form", //decide
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(normalizeCmd)
}