package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/asokawotulo/unyil/api"
	"github.com/gorilla/mux"
)

// Serve routes
func Serve() {
	// TODO: Add cors handler

	r := mux.NewRouter()

	serveAPI(r)

	s := &http.Server{
		Addr:        fmt.Sprintf(":%d", 8080),
		Handler:     r,
		ReadTimeout: 2 * time.Minute,
	}

	s.ListenAndServe()
}

func serveAPI(r *mux.Router) {
	r.HandleFunc("/", api.Index).Methods("GET")
}
