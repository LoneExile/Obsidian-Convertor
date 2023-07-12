package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Obsidian Convertor",
	Long:  `All software has versions. This is Obsidian Convertor's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nObsidian Convertor v0.1.6")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
