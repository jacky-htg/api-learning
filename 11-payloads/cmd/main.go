package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/go-services/libraries/config"
	"github.com/jacky-htg/go-services/libraries/database"
	"github.com/jacky-htg/go-services/schema"
)

func main() {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	if err := run(); err != nil {
		log.Printf("error: shutting down: %s", err)
		os.Exit(1)
	}
}

func run() error {

	flag.Parse()

	// =========================================================================
	// Start Database

	db, err := database.Open()
	if err != nil {
		return fmt.Errorf("connecting to db: %v", err)
	}
	defer db.Close()

	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			return fmt.Errorf("applying migrations: %v", err)
		}
		fmt.Println("Migrations complete")

	case "seed":
		if err := schema.Seed(db); err != nil {
			return fmt.Errorf("seeding database: %v", err)
		}
		fmt.Println("Seed data complete")
	}

	return nil
}
