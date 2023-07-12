package converter

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/LoneExile/obsidian-convertor/utils"
	"github.com/davidbyttow/govips/v2/vips"
)

type Config struct {
	InputPath       string
	ImagePath       string
	OutputPathMD    string
	OutputPathImg   string
	CustomImagePath string
	OutputFormat    string
	Quality         int
}

type Converter struct {
	Config Config
}

func NewConverter(config Config) *Converter {
	return &Converter{
		Config: config,
	}
}

func (c *Converter) ConvertObsidianToMarkdown() error {
	// Set vips logging settings
	vips.LoggingSettings(func(messageDomain string, messageLevel vips.LogLevel, message string) {}, vips.LogLevelError)

	// Begin govips conversion
	vips.Startup(nil)
	defer vips.Shutdown()

	inputPath := c.Config.InputPath
	imagePath := c.Config.ImagePath
	outputPathMD := c.Config.OutputPathMD
	outputPathImg := c.Config.OutputPathImg
	customImagePath := c.Config.CustomImagePath
	outputFormat := c.Config.OutputFormat
	quality := c.Config.Quality

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

				err = os.MkdirAll(filepath.Dir(newImagePath), os.ModePerm)
				if err != nil {
					return err
				}

				if outputFormat == "same" {
					err = utils.CopyImage(imagePathWithExt, newImagePath)
					if err != nil {
						return err
					}
				} else {
					format := utils.ImageFormat(outputFormat)
					quality := utils.ImageQuality(quality)
					err = utils.ConvertImage(imagePathWithExt, newImagePath, format, quality)
					if err != nil {
						return err
					}
				}

				relImagePath, err := filepath.Rel(filepath.Dir(newOutputPath), outputPathImg)
				if err != nil {
					return err
				}

				var newImageNameWithNewExt string
				ext := filepath.Ext(newImageName)
				newImageNameWithoutExt := strings.TrimSuffix(newImageName, ext)

				if outputFormat == "same" {
					newImageNameWithNewExt = newImageName
				} else {
					newImageNameWithNewExt = newImageNameWithoutExt + "." + outputFormat
				}

				outputImagePathMD := ""
				if customImagePath != "" {
					outputImagePathMD = filepath.Join(customImagePath, newImageNameWithNewExt)
				} else {
					outputImagePathMD = filepath.Join(relImagePath, newImageNameWithNewExt)
				}

				line = strings.Replace(line, match[0], fmt.Sprintf("![%s](%s)", newImageNameWithoutExt, outputImagePathMD), 1)
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
