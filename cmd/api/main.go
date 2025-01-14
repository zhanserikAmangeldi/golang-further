package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Variable that will save version of the application. Hard-coding
const version = "1.0.0"

// Structure to hold all the configuration settings for our application
type config struct {
	port int
	env  string
}

// Structure to hold the dependencies for our HTTP handlers, helpers and middleware. Except that there are logger.
// It will grow to include a lot more as our build progress
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// using configuration structure to get the values of port and env, to transfer to flags
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Init a new logger which writes messages to the standard out stream.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Instance of the app structure, containing struct and the logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	// New servemux and add a /1/healthcheck route which dispatches requests
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	// HTTP server with configuration
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the HTTP server using ListenAndServe()
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
