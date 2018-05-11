package logger

import (
	"log"
	"time"

	"github.com/valyala/fasthttp"
)

// Logger show request information
func Logger(innerHandler fasthttp.RequestHandler, name string) fasthttp.RequestHandler {
	return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
		start := time.Now()

		path := string(ctx.Path())

		innerHandler(ctx)

		log.Printf(
			"%s\t%s\t%s\t%s",
			string(ctx.Method()),
			path,
			name,
			time.Since(start),
		)
	})
}
