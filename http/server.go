package http

import (
	"log"

	"github.com/valyala/fasthttp"

	"github.com/efrenfuentes/imageproxy/http/routers"
	"github.com/efrenfuentes/imageproxy/http/settings"
)

// Server is our http server
type Server struct{}

// Run execute our http server
func (s *Server) Run() {
	log.Printf("Loading settings...\n")
	settings.Init()

	mySettings := settings.Get()
	ip := mySettings["server"].(map[string]interface{})["ip"].(string)
	port := mySettings["server"].(map[string]interface{})["port"].(string)

	listenIn := ip + ":" + port

	log.Printf("Creating routes...\n")
	routes := routers.Init()

	log.Printf("Server starting on port %v [%s]\n",
		listenIn,
		settings.GetEnvironment())

	log.Fatal(fasthttp.ListenAndServe(listenIn, routes.Handler))
}
