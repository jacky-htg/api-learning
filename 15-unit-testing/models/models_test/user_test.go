package models_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/go-cmp/cmp"
	"github.com/jacky-htg/go-services/models"
	"github.com/jacky-htg/go-services/schema"
	"github.com/jacky-htg/go-services/services/config"
	"github.com/jacky-htg/go-services/tests"
)

func init() {
	config.Setup("../../.env")

}

func TestUsers(t *testing.T) {
	db, teardown := tests.NewUnit(t)
	defer teardown()

	u0 := models.User{
		Username: "Aladin",
		Email:    "aladin@gmail.com",
		Password: "1234",
		IsActive: true,
	}

	err := u0.Create(db)
	if err != nil {
		t.Fatalf("creating user u0: %s", err)
	}

	u1 := models.User{
		ID: u0.ID,
	}

	err = u1.Get(db)
	if err != nil {
		t.Fatalf("getting user u1: %s", err)
	}

	if diff := cmp.Diff(u1, u0); diff != "" {
		t.Fatalf("fetched != created:\n%s", diff)
	}
}

func TestUserList(t *testing.T) {
	db, teardown := tests.NewUnit(t)
	defer teardown()

	if err := schema.Seed(db); err != nil {
		t.Fatal(err)
	}

	var user models.User
	users, err := user.List(db)
	if err != nil {
		t.Fatalf("listing users: %s", err)
	}
	if exp, got := 1, len(users); exp != got {
		t.Fatalf("expected users list size %v, got %v", exp, got)
	}
}
