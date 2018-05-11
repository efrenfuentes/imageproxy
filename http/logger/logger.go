package logger

import (
	"github.com/valyala/fasthttp"
)

func Logger(inner fasthttp.RequestHandler, name string) fasthttp.RequestHandler {
	// return fasthttp.RequestHandler(func(ctx *fasthttp.RequestCtx) {
	// 	start := time.Now()

	// 	inner.ServeHTTP(ctx)

	// 	log.Printf(
	// 		"%s\t%s\t%s\t%s",
	// 		r.Method,
	// 		r.RequestURI,
	// 		name,
	// 		time.Since(start),
	// 	)
	// })
}
