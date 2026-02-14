// CMPS3162 Lab 3 demonstrating writeJSON and readJSON.
package main

import (
	"expvar"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/andreshungbz/lab3-json-handling/internal/vcs"
)

var (
	version = vcs.Version()
)

// config stores the API server configuration.
type config struct {
	port int    // API server port
	env  string // (development|staging|production)
}

// application holds the dependencies for the HTTP handlers, helpers, middleware,
// etc. so that they are all accessible through dependency injection.
type application struct {
	config config
	logger *slog.Logger
	wg     sync.WaitGroup
}

func main() {
	var cfg config
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// FLAGS

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	displayVersion := flag.Bool("version", false, "Display program version")

	flag.Parse()

	// display program version and exit if the version flag was passed
	if *displayVersion {
		fmt.Printf("Version:\t%s\n", version)
		os.Exit(0)
	}

	// DATABASE

	// METRICS

	expvar.NewString("version").Set(version)

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	expvar.Publish("timestamp", expvar.Func(func() any {
		return time.Now().Unix()
	}))

	// APPLICATION

	app := &application{
		config: cfg,
		logger: logger,
	}

	// start the API server
	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
