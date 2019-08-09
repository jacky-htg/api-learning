package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jacky-htg/go-services/models"
	"github.com/jmoiron/sqlx"
)

// Users struct for test users
type User struct {
	Db *sqlx.DB
}

//Create : unit test  for create user function
func (u *User) Create(t *testing.T) {
	u0 := models.User{
		Username: "Aladin",
		Email:    "aladin@gmail.com",
		Password: "1234",
		IsActive: true,
	}

	err := u0.Create(u.Db)
	if err != nil {
		t.Fatalf("creating user u0: %s", err)
	}

	u1 := models.User{
		ID: u0.ID,
	}

	err = u1.Get(u.Db)
	if err != nil {
		t.Fatalf("getting user u1: %s", err)
	}

	if diff := cmp.Diff(u1, u0); diff != "" {
		t.Fatalf("fetched != created:\n%s", diff)
	}
}

//List : unit test for user list function
func (u *User) List(t *testing.T) {
	var user models.User
	users, err := user.List(u.Db)
	if err != nil {
		t.Fatalf("listing users: %s", err)
	}
	if exp, got := 1, len(users); exp != got {
		t.Fatalf("expected users list size %v, got %v", exp, got)
	}
}
