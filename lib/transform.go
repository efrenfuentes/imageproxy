package lib

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/disintegration/imaging"
	"github.com/efrenfuentes/imageproxy/http/settings"
)

type geometryStruct struct {
	width  int
	height int
}

// TransformImage do the image transformations
func TransformImage(path, srcImagePath, geometry string) (string, error) {
	mySettings := settings.Get()
	cacheDir := mySettings["images"].(map[string]interface{})["cache_dir"].(string)
	loggerTransform := mySettings["logger"].(map[string]interface{})["transform"].(string)

	geometryData, err := getGeometry(geometry)
	if err != nil {
		log.Printf("invalid geometry: %v", err)
		return "", err
	}

	imagePath := cacheDir + HashDir(geometry, geometry+"/"+path)

	// Open the source image.
	src, err := imaging.Open(srcImagePath)
	if err != nil {
		log.Printf("failed to open image: %v", err)
		return "", err
	}

	// Resize the image preserving the aspect ratio.
	image := imaging.Fill(src, geometryData.width, geometryData.height, imaging.Center, imaging.Lanczos)

	// Save the resulting image as JPEG.
	// Create the directory if not exists
	err = os.MkdirAll(filepath.Dir(imagePath), 0777)
	if err != nil {
		log.Printf("can't create cache directory: %v", err)
		return "", err
	}

	err = imaging.Save(image, imagePath)
	if err != nil {
		log.Printf("failed to save image: %v", err)
		return "", err
	}

	if loggerTransform == "on" {
		log.Printf("%s transformed to %s", srcImagePath, imagePath)
	}

	return imagePath, nil

}

func getGeometry(geometry string) (geometryStruct, error) {
	geometryRE, _ := regexp.Compile(`(\d+)x(\d+)|original`) // Prepare our regex
	result := geometryRE.FindStringSubmatch(geometry)
	if len(result) > 0 {
		width, _ := strconv.Atoi(result[1])
		height, _ := strconv.Atoi(result[2])

		return geometryStruct{width: width, height: height}, nil
	}

	return geometryStruct{}, errors.New("invalid geometry")

}
