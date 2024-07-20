package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Config instance
	var cfg config

	// Read value of port anf env from command line
	// default to 4000 and development of no corresponding flag provided
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment(development|staging|production)")
	flag.Parse()

	// initialize logger

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Application instance

	app := &application{
		config: cfg,
		logger: logger,
	}



	// Declear http server

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)

	fmt.Println("Hello world")
}
