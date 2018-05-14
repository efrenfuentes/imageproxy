package lib

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// SaveOnCache store the image on disk to avoid download it again
func SaveOnCache(imageURL string, path string) (string, error) {
	// Create the directory if not exists
	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		return "", errors.New("can't create cache directory")
	}

	// Create the destination file
	outputFile, err := os.Create(path)
	if err != nil {
		return "", errors.New("can't create cache file")
	}
	defer outputFile.Close()

	// Get the data from URL
	response, err := http.Get(imageURL)

	if err != nil || response.StatusCode != 200 {
		return "", errors.New("can't download the image")
	}
	defer response.Body.Close()

	// Write the body to file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		return "", err
	}

	return path, nil
}
