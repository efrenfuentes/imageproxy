package http

import (
	"log"
	"net/http"

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
	http.Handle("/", routes)

	log.Printf("Server starting on port %v [%s]\n",
		listenIn,
		settings.GetEnvironment())

	log.Fatal(http.ListenAndServe(listenIn, nil))
}
