package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var config config
	flag.IntVar(&config.port, "port", 4000, "Server port to listen on")
	flag.StringVar(&config.env, "env", "development", "Application environment (development|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

	app := &application{
		config: config,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", config.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
