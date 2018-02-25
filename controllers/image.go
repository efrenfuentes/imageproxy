package controllers

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"

	"github.com/efrenfuentes/imageproxy/lib"
	"github.com/gorilla/mux"
)

func ImageIndex(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	// geometry := urlParams["geometry"]
	path := urlParams["path"]

	image, format, err := lib.DownloadImage(path)

	if err != nil || format == "error" {
		http.Error(w, "Sorry can't find a valid image on the requested url!", 415)
	}

	// transform the image
	//   return error if geometry is not valid
	//   return error if can't transform image
	// return the image

	if err := writeImage(w, image, format); err != nil {
		http.Error(w, "Image can't be served", 500)
	}
}

// writeImage encodes an image 'img' in correct format and writes it into ResponseWriter.
func writeImage(w http.ResponseWriter, img image.Image, format string) error {

	buffer := new(bytes.Buffer)

	if format == "" {
		format = "png"
	}

	if format == "jpeg" {
		w.Header().Set("Content-Type", "image/jpeg")
		if err := jpeg.Encode(buffer, img, nil); err != nil {
			return err
		}
	} else if format == "png" {
		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(buffer, img); err != nil {
			return err
		}
	} else if format == "gif" {
		w.Header().Set("Content-Type", "image/gif")
		if err := gif.Encode(buffer, img, nil); err != nil {
			return err
		}
	}

	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	if _, err := w.Write(buffer.Bytes()); err != nil {
		return err
	}

	return nil
}
