package app

import (
	"github.com/gorilla/mux"
)

var Router *mux.Router

type server struct {
	*mux.Router
}

func init() {
	Router = newServer().registerRoutes()
}

func newServer() server {
	return server{
		mux.NewRouter(),
	}
}

func (s server) registerRoutes() *mux.Router {
	s.Router.HandleFunc("/", HomeHandler)
	s.Router.HandleFunc("/ws", WebSocketEndpoint)

	return s.Router
}
