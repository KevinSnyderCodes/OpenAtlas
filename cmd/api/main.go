package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/api"
	"github.com/gorilla/mux"
)

var (
	fPort    = flag.Int("port", 8080, "Port to listen on.")
	fPortSSL = flag.Int("port-ssl", 8443, "Port to listen on for SSL.")

	fSSLCertFile = flag.String("ssl-cert-file", "./cert/cert.pem", "SSL certificate file.")
	fSSLKeyFile  = flag.String("ssl-key-file", "./cert/key.pem", "SSL key file.")
)

var HandlerFuncNotImplemented = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func run() error {
	flag.Parse()

	var organizations api.Organizations = &api.DefaultOrganizations{}
	var workspaces api.Workspaces = &api.DefaultWorkspaces{}

	r := mux.NewRouter()

	r.Methods("GET").Path("/.well-known/terraform.json").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := map[string]string{
			"tfe.v2":   "/api/v2/",
			"tfe.v2.1": "/api/v2/",
		}

		data, err := json.Marshal(m)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	s := r.PathPrefix("/api/v2").Subrouter()

	s.Methods("GET").Path("/organizations/{organization}/entitlement-set").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		organization := vars["organization"]

		if organization != "default" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp, err := organizations.ReadEntitlements(r.Context(), organization)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := resp.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/vnd.api+json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	s.Methods("GET").Path("/organizations/{organization}/workspaces/{workspace}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		organization := vars["organization"]
		workspace := vars["workspace"]

		if organization != "default" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if workspace != "default" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp, err := workspaces.Read(r.Context(), organization, workspace)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data, err := resp.MarshalJSON()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/vnd.api+json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})

	go func() {
		addr := fmt.Sprintf(":%d", *fPort)
		fmt.Printf("Listening on %s\n", addr)
		if err := http.ListenAndServe(addr, r); err != nil {
			panic(fmt.Errorf("error listening and serving: %w", err))
		}
	}()

	go func() {
		addrSSL := fmt.Sprintf(":%d", *fPortSSL)
		fmt.Printf("Listening on %s\n", addrSSL)
		if err := http.ListenAndServeTLS(addrSSL, *fSSLCertFile, *fSSLKeyFile, r); err != nil {
			panic(fmt.Errorf("error listening and serving: %w", err))
		}
	}()

	ch := make(chan struct{})
	<-ch

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
