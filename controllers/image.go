package controllers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/efrenfuentes/imageproxy/lib"
	"github.com/valyala/fasthttp"
)

// ImageIndex serve the image
func ImageIndex(ctx *fasthttp.RequestCtx) {
	geometry := strings.ToLower(ctx.UserValue("geometry").(string))
	path := ctx.UserValue("path").(string)[1:]
	extension := strings.ToLower(filepath.Ext(path))

	key := path + version(ctx)
	e := `"` + key + `"`
	ctx.Response.Header.Set("Etag", e)
	ctx.Response.Header.Set("Cache-Control", "public,max-age=60400") // 7 days

	match := string(ctx.Request.Header.Peek("If-None-Match")[:])
	if match == e {
		if strings.Contains(match, e) {
			// TODO: Check if image has changes
			ctx.SetStatusCode(304)
			return
		}
	}

	if validGeometry(geometry) && validExtension(extension) {
		filepath, err := lib.DownloadImage(path) // Download the original image

		if err != nil { // we can't download the image
			err = os.Remove(filepath)
			ctx.SetStatusCode(415)
			fmt.Fprint(ctx, "Sorry can't find a valid image on the requested url!")
		} else { // we are ready to serve the image as static file
			// transform image if is needed
			if geometry != "original" {
				filepath, err = lib.TransformImage(path, filepath, geometry, version(ctx))
			}

			if err != nil {
				err = os.Remove(filepath)
				ctx.SetStatusCode(415)
				fmt.Fprint(ctx, "Sorry can't find a valid image on the requested url!")
			} else {
				fasthttp.ServeFile(ctx, filepath)
			}
		}
	} else {
		ctx.SetStatusCode(415)
		fmt.Fprint(ctx, "We only handle images!")
	}
}

func validGeometry(geometry string) bool {
	geometryRE, _ := regexp.Compile(`(\d+)x(\d+)|original`) // Prepare our regex
	result := geometryRE.FindStringSubmatch(geometry)

	if len(result) > 0 {
		return true
	}

	return false
}

func validExtension(extension string) bool {
	extensionRE, _ := regexp.Compile(`jpg|jpeg|png|gif`) // Prepare our regex
	result := extensionRE.FindStringSubmatch(extension)

	if len(result) > 0 {
		return true
	}

	return false
}

func version(ctx *fasthttp.RequestCtx) string {
	version := string(ctx.QueryArgs().Peek("ver"))

	if version != "" {
		version = "-" + version
	}

	return version
}
