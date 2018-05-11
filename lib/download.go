package lib

import (
	"log"
	"os"

	"github.com/efrenfuentes/imageproxy/http/settings"
)

// DownloadImage download the image to path
func DownloadImage(path string) error {
	mySettings := settings.Get()
	imagesEndPoint := mySettings["images"].(map[string]interface{})["endpoint"].(string)
	cacheDir := mySettings["images"].(map[string]interface{})["cache_dir"].(string)
	loggerCache := mySettings["logger"].(map[string]interface{})["cache"].(string)

	filePath := cacheDir + "original/" + path

	if _, err := os.Stat(filePath); err == nil { // File already on cache
		if loggerCache == "on" {
			log.Printf("on cache %s\n", filePath)
		}
	} else { // We need download the file
		imageURL := imagesEndPoint + path
		if loggerCache == "on" {
			log.Printf("dowloading %s\n", imageURL)
		}

		// Store image on cache
		return SaveOnCache(imageURL, filePath)
	}

	return nil
}
