package main

import (
	"github.com/samuelyang09191/web-socket/cmd/app"
	"net/http"
)

func main() {
	s := http.Server{
		Addr:    ":8080",
		Handler: app.DefaultServer.Router,
	}

	s.ListenAndServe()
}
