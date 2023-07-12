package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "obsidian-convertor",
	Short: "A converter for Obsidian markdown files",
	Long:  `A command-line tool that converts Obsidian markdown files to regular markdown files`,
}

func Execute() *cobra.Command {
	rCmd, err := rootCmd.ExecuteC()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return rCmd
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// define any configuration setup
	// https://github.com/spf13/viper
	// viper.SetConfigName("config")
}
