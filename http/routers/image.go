package routers

import (
	"github.com/efrenfuentes/imageproxy/controllers"
	"github.com/gorilla/mux"
)

func SetImageRoutes(router *mux.Router) *mux.Router {

	SetRoute(router, "/{geometry}/{path:.+}", "ImageIndex", "GET", controllers.ImageIndex)

	return router
}
