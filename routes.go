package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aaron2198/vts_broker/handler"
	repo "github.com/aaron2198/vts_broker/repository"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// routes Core application router, handles all requests on the environments listening port
func routes(e *Env) {
	var wait time.Duration
	router := mux.NewRouter()

	// Bootstrap all known routes
	e.GetRoutes(router)
	// Bootstrap all known middleware
	e.GetMiddleware(router)

	// Construct the http server
	port := fmt.Sprintf("%d", e.Conf.port)
	srv := &http.Server{
		Addr: "0.0.0.0:" + port,
		// To avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	e.Logger.Log.WithFields(logrus.Fields{
		"port": port,
	}).Info("API service started")

	// create a goroutine handle listen and serve without blocking requests
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			e.Logger.Log.Error(err)
		}
	}()

	// We accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		e.Logger.Log.Error("Error trying to shutdown: %e", err)
	}
	e.Logger.Log.Info("Shutting down")
	os.Exit(0)

}

func (e *Env) GetRoutes(router *mux.Router) {
	router.HandleFunc("/hc", e.healthHandler).Methods("GET")
	router.HandleFunc("/communities", e.community_route)
	router.HandleFunc("/instancedb", e.instanceDb_route)
}

func (e *Env) GetMiddleware(router *mux.Router) {
	router.Use(e.loggingMW)
	router.Use(corsMW)
	// Only logs on trace even without this, but it skips some work avoiding the call
	if e.Conf.logLevel == logrus.TraceLevel {
		router.Use(e.profilerMW)
	}
}

// Routes

// healthHandler /hc [GET]
func (e *Env) healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// communitiesHandler CRUD communities
func (e *Env) community_route(w http.ResponseWriter, r *http.Request) {
	// Choose and construct an instance of our data access object (DAO)
	// This repo.Community is a struct that implements vtsb_interface.Community
	// Provide the repo the environments db connection.
	repository := &repo.Community{
		Db: e.Db,
	}
	// Here we create a handler, this wires our data access object (repository) to
	// the type of incoming request, we provide our writer to this handler as it will
	// translate the result of our repository to a format for the end user ie. json.
	h := &handler.Community{
		CommunityInterface: repository,
		Logger:             e.Logger,
	}
	h.Handle(r, w)
}

func (e *Env) instanceDb_route(w http.ResponseWriter, r *http.Request) {
	// Choose and construct an instance of our data access object (DAO)
	repository := &repo.InstanceDb{
		Db: e.Db,
	}
	h := &handler.InstanceDb{
		InstanceDbInterface: repository,
		Logger:              e.Logger,
	}
	h.Handle(r, w)
}
