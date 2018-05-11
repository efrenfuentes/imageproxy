package controllers

import (
	"fmt"
	"strings"

	"github.com/efrenfuentes/imageproxy/http/settings"
	"github.com/efrenfuentes/imageproxy/lib"
	"github.com/valyala/fasthttp"
)

// ImageIndex serve the image
func ImageIndex(ctx *fasthttp.RequestCtx) {
	mySettings := settings.Get()
	cacheDir := mySettings["images"].(map[string]interface{})["cache_dir"].(string)

	geometry := strings.ToLower(ctx.UserValue("geometry").(string))
	path := ctx.UserValue("path").(string)[1:]

	err := lib.DownloadImage(path) // Download the original image

	// transform image if is needed
	if geometry != "original" {
		lib.TransformImage(path, geometry)
	}

	if err != nil { // we can't download the image
		ctx.SetStatusCode(415)
		fmt.Fprint(ctx, "Sorry can't find a valid image on the requested url!")
	} else { // we are ready to serve the image as static file
		fasthttp.ServeFile(ctx, cacheDir+string(ctx.Path()))
	}
}
