package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/KevinSnyderCodes/OpenAtlas/internal/api"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/db"
	"github.com/KevinSnyderCodes/OpenAtlas/internal/tasks"
	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"

	"github.com/jackc/pgx/v5"
)

var (
	fDatabaseURL = flag.String("database-url", os.Getenv("DATABASE_URL"), "")
	fRedisURL    = flag.String("redis-url", os.Getenv("REDIS_URL"), "")

	fPort    = flag.Int("port", 8080, "")
	fPortSSL = flag.Int("port-ssl", 8443, "")

	fSSLCertFile = flag.String("ssl-cert-file", "./cert/cert.pem", "")
	fSSLKeyFile  = flag.String("ssl-key-file", "./cert/key.pem", "")
)

var HandlerFuncNotImplemented = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

var HandlerFuncMarshalJSON = func(w http.ResponseWriter, r *http.Request, resp json.Marshaler, err error) {
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
}

func run() error {
	flag.Parse()

	if *fDatabaseURL == "" {
		return fmt.Errorf("database url is required")
	}
	if *fRedisURL == "" {
		return fmt.Errorf("redis url is required")
	}

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, *fDatabaseURL)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	queries := db.New(conn)

	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: *fRedisURL})
	defer asynqClient.Close()

	asynqHandler := tasks.NewHandler(queries)

	asyncServer := asynq.NewServer(asynq.RedisClientOpt{Addr: *fRedisURL}, asynq.Config{})
	asyncMux := asynq.NewServeMux()
	asyncMux.HandleFunc(tasks.TypeRunProcess, asynqHandler.HandleRunProcessTask)

	go func() {
		if err := asyncServer.Run(asyncMux); err != nil {
			panic(fmt.Errorf("error running async server: %w", err))
		}
	}()

	fmt.Println("Async server running")

	var configurationversions api.ConfigurationVersions = api.NewDefaultConfigurationVersions(queries)
	var organizations api.Organizations = api.NewDefaultOrganization(queries)
	var plans api.Plans = api.NewDefaultPlans(queries)
	var runs api.Runs = api.NewDefaultRuns(queries)
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

	sAPI := r.PathPrefix("/api").Subrouter()

	sAPI.Methods(http.MethodPut).Path("/upload/{configuration_version_id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		configurationVersionID := vars["configuration_version_id"]

		if err := configurationversions.Upload(r.Context(), configurationVersionID, r.Body); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	sTFE := r.PathPrefix("/api/v2").Subrouter()

	sTFE.Methods(http.MethodGet).Path("/configuration-versions/{configuration_version_id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		configurationVersionID := vars["configuration_version_id"]

		resp, err := configurationversions.Read(r.Context(), configurationVersionID)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodGet).Path("/plans/{plan_id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		planID := vars["plan_id"]

		resp, err := plans.Read(r.Context(), planID)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodGet).Path("/organizations/{organization}/entitlement-set").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		organization := vars["organization"]

		resp, err := organizations.ReadEntitlements(r.Context(), organization)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodGet).Path("/organizations/{organization}/runs/queue").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		organization := vars["organization"]

		resp, err := organizations.ReadRunQueue(r.Context(), organization)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodGet).Path("/organizations/{organization}/workspaces/{workspace}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		organization := vars["organization"]
		workspace := vars["workspace"]

		resp, err := workspaces.Read(r.Context(), organization, workspace)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodPost).Path("/runs").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var req api.RunCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := runs.Create(r.Context(), &req)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		task, err := tasks.NewRunProcessTask(resp.Data.ID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if _, err := asynqClient.Enqueue(task); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodGet).Path("/runs/{run_id}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		runID := vars["run_id"]

		resp, err := runs.Read(r.Context(), runID)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodPost).Path("/workspaces/{workspace_id}/configuration-versions").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceID := vars["workspace_id"]

		var req api.ConfigurationVersionCreateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		resp, err := configurationversions.Create(r.Context(), workspaceID, &req)
		HandlerFuncMarshalJSON(w, r, resp, err)
	})

	sTFE.Methods(http.MethodGet).Path("/workspaces/{workspace_id}/runs").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		workspaceID := vars["workspace_id"]

		if workspaceID != "ws-0000000000000000" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		resp, err := runs.List(r.Context(), workspaceID)
		HandlerFuncMarshalJSON(w, r, resp, err)
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
