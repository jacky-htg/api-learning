package controllers

import (
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/libraries/database"
	"github.com/jmoiron/sqlx"
)

//Checks : struct for set Checks Dependency Injection
type Checks struct {
	Db *sqlx.DB
}

//Login : http handler for login
func (u *Checks) Health(w http.ResponseWriter, r *http.Request) error {
	var health struct {
		Status string `json:"status"`
	}

	// Check if the database is ready.
	if err := database.StatusCheck(r.Context(), u.Db); err != nil {

		// If the database is not ready we will tell the client and use a 500
		// status. Do not respond by just returning an error because further up in
		// the call stack will interpret that as an unhandled error.
		health.Status = "db not ready"
		return err
	}

	health.Status = "ok"
	return api.ResponseOK(w, health, http.StatusOK)
}
