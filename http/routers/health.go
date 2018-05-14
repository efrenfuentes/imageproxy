package routers

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/efrenfuentes/imageproxy/controllers"
)

// SetHealthRoutes set routes for images
func SetHealthRoutes(router *fasthttprouter.Router) *fasthttprouter.Router {
	SetRoute(router, "/", "HealthIndex", "GET", controllers.HealthIndex)

	return router
}
