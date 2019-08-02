package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jacky-htg/go-services/payloads/response"
	"github.com/jmoiron/sqlx"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/models"
)

//Users : struct for set Users Dependency Injection
type Users struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	var user models.User
	list, err := user.List(u.Db)
	if err != nil {
		u.Log.Printf("error call list users: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var listResponse []*response.UserResponse
	for _, user := range list {
		var userResponse response.UserResponse
		userResponse.Transform(user)
		listResponse = append(listResponse, &userResponse)
	}

	data, err := json.Marshal(listResponse)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}

//View : http handler for retrieve user by id
func (u *Users) View(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user models.User
	user.ID = uint64(id)
	err = user.Get(u.Db)
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var response response.UserResponse
	response.Transform(user)
	data, err := json.Marshal(response)
	if err != nil {
		u.Log.Println("error marshalling result", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		u.Log.Println("error writing result", err)
	}
}
