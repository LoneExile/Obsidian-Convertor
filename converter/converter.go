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

	cc := c.Config

	err := filepath.Walk(cc.InputPath, func(path string, info os.FileInfo, err error) error {
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

		relPath, _ := filepath.Rel(cc.InputPath, path)
		newOutputPath := filepath.Join(cc.OutputPathMD, relPath)

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

		// TODO: scanner.Scan() each line is slow?, maybe read until EOF? or found match regex?
		for scanner.Scan() {
			line := scanner.Text()
			matches := re.FindAllStringSubmatch(line, -1)

			for _, match := range matches {
				imageName := match[1]
				imagePathWithExt := filepath.Join(cc.ImagePath, imageName)
				if _, err := os.Stat(imagePathWithExt); os.IsNotExist(err) {
					return fmt.Errorf("image not found: %s", imagePathWithExt)
				}

				newImageName := strings.ReplaceAll(imageName, " ", "_")
				newImagePath := filepath.Join(cc.OutputPathImg, newImageName)

				err = os.MkdirAll(filepath.Dir(newImagePath), os.ModePerm)
				if err != nil {
					return err
				}

				if cc.OutputFormat == "same" {
					err = utils.CopyImage(imagePathWithExt, newImagePath)
					if err != nil {
						return err
					}
				} else {
					format := utils.ImageFormat(cc.OutputFormat)
					quality := utils.ImageQuality(cc.Quality)
					err = utils.ConvertImage(imagePathWithExt, newImagePath, format, quality)
					if err != nil {
						return err
					}
				}

				relImagePath, err := filepath.Rel(filepath.Dir(newOutputPath), cc.OutputPathImg)
				if err != nil {
					return err
				}

				var newImageNameWithNewExt string
				ext := filepath.Ext(newImageName)
				newImageNameWithoutExt := strings.TrimSuffix(newImageName, ext)

				if cc.OutputFormat == "same" {
					newImageNameWithNewExt = newImageName
				} else {
					newImageNameWithNewExt = newImageNameWithoutExt + "." + cc.OutputFormat
				}

				outputImagePathMD := ""
				if cc.CustomImagePath != "" {
					outputImagePathMD = filepath.Join(cc.CustomImagePath, newImageNameWithNewExt)
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
