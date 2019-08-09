package main

import (
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	userUnitTest "github.com/jacky-htg/go-services/models/tests"
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

	// unit test for user
	user := userUnitTest.User{Db: db}
	t.Run("UsersList", user.List)
	t.Run("UsersCreate", user.Create)
}
