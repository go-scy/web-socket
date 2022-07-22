package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
	s.Router.HandleFunc("/home", HomeHandler)
	s.Router.HandleFunc("/ws", WebSocketEndpoint)

	return s.Router
}
