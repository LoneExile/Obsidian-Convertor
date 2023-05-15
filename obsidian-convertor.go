package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "obsidian-convertor",
	Short: "A converter for Obsidian markdown files",
	Long:  `A command-line tool that converts Obsidian markdown files to regular markdown files`,
}

var convertCmd = &cobra.Command{
	Use:   "convert <input-path> <image-path> <output-path> <output-image-path> [<custom-image-path>]",
	Short: "Converts Obsidian markdown files to regular markdown files",
	Long:  `This command converts Obsidian markdown files to regular markdown files.`,
	Args:  cobra.RangeArgs(4, 5), // This line enforces that the command receives between 4 and 5 arguments
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\n Obsidian Convertor")
		fmt.Println("-------------------")

		inputPath := args[0]
		imagePath := args[1]
		outputPathMD := args[2]
		outputPathImg := args[3]

		customImagePath := ""
		if len(args) == 5 {
			customImagePath = args[4]
		}

		fmt.Println("Input path:", inputPath)
		fmt.Println("Image path:", imagePath)
		fmt.Println("Output path (markdown):", outputPathMD)
		fmt.Println("Output path (images):", outputPathImg)
		fmt.Println("Custom image path:", customImagePath)

		if err := convertObsidianToMarkdown(inputPath, imagePath, outputPathMD, outputPathImg, customImagePath); err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println("-------------------")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Obsidian Convertor",
	Long:  `All software has versions. This is Obsidian Convertor's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nObsidian Convertor v0.1.3") // Change the version number as needed
	},
}

func convertObsidianToMarkdown(inputPath, imagePath, outputPathMD, outputPathImg, customImagePath string) error {
	err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".md" {
			return nil
		}

		inputFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer inputFile.Close()

		// NOTE: this is for copying the file to the output directory w/o copying the directory structure
		// outputFileName := filepath.Base(path)
		// outputFile, err := os.Create(filepath.Join(outputPathMD, outputFileName))

		relPath, _ := filepath.Rel(inputPath, path)
		newOutputPath := filepath.Join(outputPathMD, relPath)

		newOutputDir := filepath.Dir(newOutputPath)
		if err := os.MkdirAll(newOutputDir, os.ModePerm); err != nil {
			return err
		}
		outputFile, err := os.Create(newOutputPath)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		re := regexp.MustCompile(`!\[\[(.*?)\]\]`)
		scanner := bufio.NewScanner(inputFile)
		writer := bufio.NewWriter(outputFile)

		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindAllStringSubmatch(line, -1)

			for _, match := range matches {
				imageName := match[1]
				imagePathWithExt := filepath.Join(imagePath, imageName)
				if _, err := os.Stat(imagePathWithExt); os.IsNotExist(err) {
					return fmt.Errorf("image not found: %s", imagePathWithExt)
				}

				number := regexp.MustCompile(`\d+`).FindString(imageName)
				newImageName := fmt.Sprintf("%s%s", number, imageName)
				newImagePath := filepath.Join(outputPathImg, newImageName)

				srcImage, err := os.Open(imagePathWithExt)
				if err != nil {
					return err
				}
				defer srcImage.Close()

				err = os.MkdirAll(filepath.Dir(newImagePath), os.ModePerm)
				if err != nil {
					return err
				}

				dstImage, err := os.Create(newImagePath)
				if err != nil {
					return err
				}
				defer dstImage.Close()

				if _, err := io.Copy(dstImage, srcImage); err != nil {
					return err
				}

				relImagePath, err := filepath.Rel(filepath.Dir(newOutputPath), outputPathImg)
				if err != nil {
					return err
				}

				outputImagePathMD := ""
				if customImagePath != "" {
					outputImagePathMD = filepath.Join(customImagePath, newImageName)
				} else {
					outputImagePathMD = filepath.Join(relImagePath, newImageName)
				}

				line = strings.Replace(line, match[0], fmt.Sprintf("![%s](%s)", imageName, outputImagePathMD), 1)
			}

			if _, err := writer.WriteString(line + "\n"); err != nil {
				return err
			}
		}

		if err := writer.Flush(); err != nil {
			return err
		}

		return nil
	})

	return err
}

func main() {
	rootCmd.AddCommand(convertCmd)
	rootCmd.AddCommand(versionCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
