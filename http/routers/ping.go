package routers

import (
	"github.com/efrenfuentes/imageproxy/controllers"
	"github.com/gorilla/mux"
)

func SetPingRoutes(router *mux.Router) *mux.Router {

	SetRoute(router, "/ping", "PingIndex", "GET", controllers.PingIndex)

	return router
}
