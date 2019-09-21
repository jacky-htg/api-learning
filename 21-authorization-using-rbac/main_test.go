package main

import (
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jacky-htg/go-services/libraries/config"
	apiTest "github.com/jacky-htg/go-services/packages/auth/controllers/tests"
	"github.com/jacky-htg/go-services/routing"
	"github.com/jacky-htg/go-services/schema"
	"github.com/jacky-htg/go-services/tests"
)

var token string

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

	// api test for auths
	{
		auths := apiTest.Auths{App: routing.API(db, log)}
		t.Run("ApiLogin", auths.Login)
		token = auths.Token
	}

	// api test for users
	{
		users := apiTest.Users{App: routing.API(db, log)}
		t.Run("APiUsersList", users.List)
		t.Run("APiUsersCrud", users.Crud)
	}
}
