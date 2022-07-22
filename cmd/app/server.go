package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	*mux.Router
}

var DefaultServer server

func init() {
	DefaultServer.Router = mux.NewRouter()
	DefaultServer.registerRoutes()
}

func (s server) registerRoutes() {
	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, http.StatusOK)
	})
	s.Router.HandleFunc("/home", HomeHandler)
}
