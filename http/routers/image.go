package routers

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/efrenfuentes/imageproxy/controllers"
)

// SetImageRoutes set routes for images
func SetImageRoutes(router *fasthttprouter.Router) *fasthttprouter.Router {
	SetRoute(router, "/:geometry/*path", "ImageIndex", "GET", controllers.ImageIndex)

	return router
}
