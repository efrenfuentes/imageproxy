package lib

import (
	"fmt"
	"image"
	"net/http"
	"os"

	"github.com/efrenfuentes/imageproxy/http/settings"
)

// DownloadImage download the image to path
func DownloadImage(path string) (image.Image, string, error) {
	var format string
	var myImage image.Image

	mySettings := settings.Get()
	imagesEndPoint := mySettings["images"].(map[string]interface{})["endpoint"].(string)
	cacheDir := mySettings["images"].(map[string]interface{})["cache_dir"].(string)
	loggerCache := mySettings["logger"].(map[string]interface{})["cache"].(string)

	cacheEnable := false
	if cacheDir != "" {
		cacheEnable = true
	}

	filePath := cacheDir + path

	if _, err := os.Stat(filePath); err == nil {
		if loggerCache == "on" {
			fmt.Printf("on cache %s\n", filePath)
		}

		infile, err := os.Open(filePath)
		if err != nil {
			return nil, "error", err
		}
		defer infile.Close()

		// Decode will figure out what type of image is in the file on its own.
		// We just have to be sure all the image packages we want are imported.
		myImage, format, err = image.Decode(infile)
		if err != nil {
			return nil, "error", err
		}
	} else {
		imageURL := imagesEndPoint + path
		if loggerCache == "on" {
			fmt.Printf("dowloading %s\n", imageURL)
		}

		r, err := http.Get(imageURL)

		if err != nil || r.StatusCode != 200 {
			return nil, "error", err
		}

		defer r.Body.Close()
		myImage, format, err = image.Decode(r.Body)
		if err != nil {
			return nil, "error", err
		}

		if cacheEnable {
			SaveOnCache(myImage, format, filePath)
		}
	}

	return myImage, format, nil
}
