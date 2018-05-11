package lib

import (
	"errors"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

// SaveOnCache store the image on disk to avoid download it again
func SaveOnCache(img image.Image, format string, path string) error {
	// create directories if don't exist
	// save the image
	// outputFile is a File type which satisfies Writer interface
	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return errors.New("can't create cache directory")
	}

	outputFile, err := os.Create(path)
	if err != nil {
		return errors.New("can't create cache file")
	}

	// Encode takes a writer interface and an image interface
	// We pass it the File and the RGBA
	jpeg.Encode(outputFile, img, nil)

	// Don't forget to close files
	outputFile.Close()

	return nil
}
