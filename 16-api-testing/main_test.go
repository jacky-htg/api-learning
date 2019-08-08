package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	userTest "github.com/jacky-htg/go-services/controllers/tests"
	"github.com/jacky-htg/go-services/routing"
	"github.com/jacky-htg/go-services/schema"
	"github.com/jacky-htg/go-services/services/config"
	"github.com/jacky-htg/go-services/tests"
)

func TestMain(t *testing.T) {
	_, ok := os.LookupEnv("APP_ENV")
	if !ok {
		config.Setup(".env")
	}

	db, teardown := tests.NewUnit(t)
	defer teardown()

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	log := log.New(os.Stderr, "TEST : ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	users := userTest.Users{App: routing.API(db, log)}

	t.Run("UsersList", users.List)
	t.Run("UsersView", users.View)
	t.Run("UsersCreate", users.Create)
	t.Run("UsersUpdate", users.Update)
	t.Run("UsersDelete", users.Delete)
}
