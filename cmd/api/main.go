package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var fPort = flag.Int("port", 8080, "Port to listen on.")

var HandlerFuncNotImplemented = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func run() error {
	flag.Parse()

	addr := fmt.Sprintf(":%d", *fPort)

	r := mux.NewRouter()

	s := r.PathPrefix("/api/v2").Subrouter()

	// s.Methods("GET").Path("/organizations").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("GET").Path("/organizations").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("POST").Path("/organizations").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("GET").Path("/organizations/{organization}").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("PATCH").Path("/organizations/{organization}").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("DELETE").Path("/organizations/{organization}").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("GET").Path("/organizations/{organization}/capacity").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("GET").Path("/organizations/{organization}/entitlement-set").HandlerFunc(HandlerFuncNotImplemented)
	s.Methods("GET").Path("/organizations/{organization}/runs/queue").HandlerFunc(HandlerFuncNotImplemented)

	fmt.Printf("Listening on %s\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		return fmt.Errorf("error listening and serving: %w", err)
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
