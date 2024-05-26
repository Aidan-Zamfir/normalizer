package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use: "norm",
	Short: "Normalize csv data for ML",
	Long: "This tool lets you enter a csv file and returns normalized/standardized data for ML",
}

// executes root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//define flags and config settings
}