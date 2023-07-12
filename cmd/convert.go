package cmd

import (
	"fmt"
	"os"

	"github.com/LoneExile/obsidian-convertor/converter"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	quality      int
)

var convertCmd = &cobra.Command{

	Use:   "convert <input-path> <image-path> <output-path> <output-image-path> [<custom-image-path>]",
	Short: "Converts Obsidian markdown files to regular markdown files",
	Long:  `This command converts Obsidian markdown files to regular markdown files.`,
	Args:  cobra.RangeArgs(4, 5),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n Obsidian Convertor")
		fmt.Println("-------------------")

		customImagePath := ""
		if len(args) == 5 {
			customImagePath = args[4]
		}

		config := converter.Config{
			InputPath:       args[0],
			ImagePath:       args[1],
			OutputPathMD:    args[2],
			OutputPathImg:   args[3],
			CustomImagePath: customImagePath,
			OutputFormat:    outputFormat,
			Quality:         quality,
		}

		fmt.Println("Input path:", config.InputPath)
		fmt.Println("Image path:", config.ImagePath)
		fmt.Println("Output path (markdown):", config.OutputPathMD)
		fmt.Println("Output path (images):", config.OutputPathImg)
		fmt.Println("Custom image path:", customImagePath)
		fmt.Println("Output format:", outputFormat)
		fmt.Println("Quality:", quality)

		c := converter.NewConverter(config)

		if err := c.ConvertObsidianToMarkdown(); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("-------------------")
	},
}

func init() {
	convertCmd.Flags().StringVarP(&outputFormat, "format", "f", "same", "Output image format (options: jpg, png, avif, same)")
	convertCmd.Flags().IntVarP(&quality, "quality", "q", 100, "Quality for output image (1-100)")
	rootCmd.AddCommand(convertCmd)
}
