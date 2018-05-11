package routers

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	"github.com/efrenfuentes/imageproxy/http/logger"
	"github.com/efrenfuentes/imageproxy/http/settings"
)

// SetRoute create a new router for our server
func SetRoute(router *fasthttprouter.Router, path, name, method string, handlerFunc func(ctx *fasthttp.RequestCtx)) *fasthttprouter.Router {
	var handler fasthttp.RequestHandler

	mySettings := settings.Get()
	loggerRoutes := mySettings["logger"].(map[string]interface{})["routes"].(string)

	handler = handlerFunc

	if loggerRoutes == "on" {
		handler = logger.Logger(handler, name)
	}

	router.Handle(method, path, handler)

	return router
}

// Init setup the routes
func Init() *fasthttprouter.Router {
	router := fasthttprouter.New()
	router = SetImageRoutes(router)
	return router
}
