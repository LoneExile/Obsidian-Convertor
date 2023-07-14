package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/davidbyttow/govips/v2/vips"
)

type ImageFormat string

const (
	FormatJPG  ImageFormat = "jpg"
	FormatPNG  ImageFormat = "png"
	FormatAVIF ImageFormat = "avif"
	FormatWEBP ImageFormat = "webp"
	FormatSame ImageFormat = "same"
)

type ImageQuality int

// TODO: cmd: Add support for quality levels?
const (
	QualityHigh   ImageQuality = 100
	QualityMedium ImageQuality = 75
	QualityLow    ImageQuality = 50
)

// CopyImage copies an image from srcPath to dstPath.
func CopyImage(srcPath, dstPath string) error {
	dstImage, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dstImage.Close()

	srcImage, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer srcImage.Close()

	if _, err := io.Copy(dstImage, srcImage); err != nil {
		return err
	}

	return nil
}

func ConvertImage(srcPath, dstPath string, format ImageFormat, quality ImageQuality) error {
	srcImage, err := vips.NewImageFromFile(srcPath)
	if err != nil {
		return err
	}

	var buf []byte
	switch format {
	case FormatJPG:
		buf, _, err = srcImage.ExportJpeg(&vips.JpegExportParams{
			Quality: int(quality),
		})
	case FormatPNG:
		buf, _, err = srcImage.ExportPng(&vips.PngExportParams{
			Quality: int(quality),
		})
	case FormatAVIF:
		buf, _, err = srcImage.ExportAvif(&vips.AvifExportParams{
			Quality: int(quality),
		})
	case FormatWEBP:
		buf, _, err = srcImage.ExportWebp(&vips.WebpExportParams{
			Quality: int(quality),
		})
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return err
	}

	// Remove the old extension from the image path
	baseName := filepath.Base(dstPath)
	ext := filepath.Ext(dstPath)
	baseNameWithoutExt := strings.TrimSuffix(baseName, ext)
	dstPathWithoutExt := filepath.Join(filepath.Dir(dstPath), baseNameWithoutExt)

	err = os.WriteFile(dstPathWithoutExt+"."+string(format), buf, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
