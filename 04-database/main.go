package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/go-services/schema"
	"github.com/jmoiron/sqlx"
)

func main() {
	// =========================================================================
	// App Starting

	log.Printf("main : Started")
	defer log.Println("main : Completed")

	// =========================================================================
	// Start Database

	dbx, err := openDBx()
	if err != nil {
		log.Fatalf("error: connecting to db: %s", err)
	}
	defer dbx.Close()

	flag.Parse()

	switch flag.Arg(0) {
	case "migrate":
		db, err := openDB()
		if err != nil {
			log.Fatalf("error: connecting to db: %s", err)
		}
		defer dbx.Close()
		if err := schema.Migrate(db); err != nil {
			log.Println("error applying migrations", err)
			os.Exit(1)
		}
		log.Println("Migrations complete")
		return

	case "seed":
		if err := schema.Seed(dbx); err != nil {
			log.Println("error seeding database", err)
			os.Exit(1)
		}
		log.Println("Seed data complete")
		return
	}

	service := Users{Db: dbx}

	// =========================================================================
	// Start API Service

	server := http.Server{
		Addr:         "localhost:9000",
		Handler:      http.HandlerFunc(service.List),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Println("server listening on", server.Addr)
		serverErrors <- server.ListenAndServe()
	}()

	// Make a channel to listen for an interrupt or terminate signal from the OS.
	// Use a buffered channel because the signal package requires it.
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// =========================================================================
	// Shutdown

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		log.Fatalf("error: listening and serving: %s", err)

	case <-shutdown:
		log.Println("caught signal, shutting down")

		// Give outstanding requests 5 seconds to complete.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("error: gracefully shutting down server: %s", err)
			if err := server.Close(); err != nil {
				log.Printf("error: closing server: %s", err)
			}
		}
	}

	log.Println("done")
}

func openDBx() (*sqlx.DB, error) {
	return sqlx.Open("mysql", "root:pass@tcp(localhost:3306)/go-services?parseTime=true")
}

func openDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:pass@tcp(localhost:3306)/go-services?parseTime=true")
}

//User : struct of User
type User struct {
	ID       uint   `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	IsActive bool   `db:"is_active"`
}

//Users : struct for set Users Dependency Injection
type Users struct {
	Db *sqlx.DB
}

//List : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	list := []User{}
	const q = `SELECT id, username, password, email, is_active FROM users`

	err := u.Db.Select(&list, q)
	if err != nil {
		log.Printf("error: query selecting users: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		log.Println("error writing result", err)
	}
}
