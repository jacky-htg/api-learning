package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	userApiTest "github.com/jacky-htg/go-services/controllers/tests"
	"github.com/jacky-htg/go-services/libraries/config"
	userUnitTest "github.com/jacky-htg/go-services/models/tests"
	"github.com/jacky-htg/go-services/routing"
	"github.com/jacky-htg/go-services/schema"
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

	// unit test for user
	user := userUnitTest.User{Db: db}
	t.Run("UnitUsersList", user.List)
	t.Run("UnitUsersCrud", user.Crud)

	// api test for users
	users := userApiTest.Users{App: routing.API(db, log)}

	t.Run("ApiUsersList", users.List)
	t.Run("APiUsersCrud", users.Crud)
}
