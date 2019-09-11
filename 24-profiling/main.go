package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof" // Register the pprof handlers
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/go-services/libraries/config"
	"github.com/jacky-htg/go-services/libraries/database"
	"github.com/jacky-htg/go-services/routing"
	"github.com/pkg/errors"
)

func main() {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	if err := run(); err != nil {
		log.Println("shutting down", "error:", err)
		os.Exit(1)
	}
}

func run() error {

	// =========================================================================
	// Logging
	log := log.New(os.Stdout, "AUTH : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// =========================================================================
	// App Starting

	log.Println("main : Started")
	defer log.Println("main : Completed")

	// =========================================================================
	// Start Database

	db, err := database.Openx()
	if err != nil {
		return errors.Wrap(err, "connecting to db")
	}
	defer db.Close()

	// =========================================================================
	// Start API Service

	server := http.Server{
		Addr:         "localhost:" + os.Getenv("APP_PORT"),
		Handler:      routing.API(db, log),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("main: server listening on %s", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Pprof server. Beware to open this code! Don’t leave your debugging information open to the world!
	/*go func() {
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()*/

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "Starting server")

	case <-shutdown:

		log.Println("main: start shutdown")

		// Give outstanding requests 5 seconds to complete.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := server.Shutdown(ctx)
		if err != nil {
			log.Printf("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = server.Close()
		}

		if err != nil {
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}
