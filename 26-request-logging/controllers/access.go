package controllers

import (
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/models"
	"github.com/jacky-htg/go-services/payloads/response"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

//Users : struct for set Users Dependency Injection
type Access struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of access
func (u *Access) List(w http.ResponseWriter, r *http.Request) error {
	var access models.Access
	tx := u.Db.MustBegin()
	list, err := access.List(r.Context(), tx)
	if err != nil {
		tx.Rollback()
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "getting access list")
	}

	var listResponse []*response.AccessResponse
	for _, a := range list {
		var accessResponse response.AccessResponse
		accessResponse.Transform(&a)
		listResponse = append(listResponse, &accessResponse)
	}

	return api.ResponseOK(r.Context(), w, listResponse, http.StatusOK)
}
