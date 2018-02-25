package routers

import (
	"net/http"

	"github.com/efrenfuentes/imageproxy/http/logger"
	"github.com/efrenfuentes/imageproxy/http/settings"
	"github.com/gorilla/mux"
)

func SetRoute(router *mux.Router, path, name, method string, handlerFunc func(w http.ResponseWriter, r *http.Request)) *mux.Router {
	var handler http.HandlerFunc

	mySettings := settings.Get()
	loggerRoutes := mySettings["logger"].(map[string]interface{})["routes"].(string)

	handler = handlerFunc

	if loggerRoutes == "on" {
		handler = logger.Logger(handler, name)
	}

	router.HandleFunc(path, handler).Methods(method)

	return router
}

func Init() *mux.Router {
	router := mux.NewRouter()
	router = SetPingRoutes(router)
	return router
}
