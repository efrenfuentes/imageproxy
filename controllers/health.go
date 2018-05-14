package controllers

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// HealthIndex serve the image
func HealthIndex(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Healthy and running!")
}
