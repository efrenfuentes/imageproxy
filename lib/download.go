package lib

import (
	"errors"
	"fmt"
	"image"
	"net/http"
	"os"

	"github.com/efrenfuentes/imageproxy/http/settings"
)

// DownloadImage download the image to path
func DownloadImage(path string) error {
	var format string
	var myImage image.Image

	mySettings := settings.Get()
	imagesEndPoint := mySettings["images"].(map[string]interface{})["endpoint"].(string)
	cacheDir := mySettings["images"].(map[string]interface{})["cache_dir"].(string)
	loggerCache := mySettings["logger"].(map[string]interface{})["cache"].(string)

	filePath := cacheDir + "original/" + path

	if _, err := os.Stat(filePath); err == nil {
		if loggerCache == "on" {
			fmt.Printf("on cache %s\n", filePath)
		}
	} else {
		imageURL := imagesEndPoint + path
		if loggerCache == "on" {
			fmt.Printf("dowloading %s\n", imageURL)
		}

		r, err := http.Get(imageURL)

		if err != nil || r.StatusCode != 200 {
			return errors.New("can't download the image")
		}

		defer r.Body.Close()
		myImage, format, err = image.Decode(r.Body)
		if err != nil {
			return err
		}

		SaveOnCache(myImage, format, filePath)
	}

	return nil
}
