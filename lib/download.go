package lib

import (
	"image"
	"net/http"

	"github.com/efrenfuentes/imageproxy/http/settings"
)

func DownloadImage(path string) (image.Image, string, error) {
	mySettings := settings.Get()
	imagesEndPoint := mySettings["images"].(map[string]interface{})["endpoint"].(string)

	imageUrl := imagesEndPoint + path

	r, err := http.Get(imageUrl)

	if err != nil || r.StatusCode != 200 {
		return nil, "error", err
	}

	defer r.Body.Close()
	myImage, format, err := image.Decode(r.Body)
	if err != nil {
		return nil, "error", err
	}

	return myImage, format, nil
}
