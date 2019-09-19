package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/packages/auth/controllers/request"
	"github.com/jacky-htg/go-services/packages/auth/controllers/response"
	"github.com/jacky-htg/go-services/packages/auth/models"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//Users : struct for set Users Dependency Injection
type Users struct {
	Db  *sql.DB
	Log *log.Logger
}

//List : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) error {
	var user models.User
	list, err := user.List(u.Db)
	if err != nil {
		return fmt.Errorf("getting users list: %v", err)
	}

	var listResponse []*response.UserResponse
	for _, user := range list {
		var userResponse response.UserResponse
		userResponse.Transform(&user)
		listResponse = append(listResponse, &userResponse)
	}

	return api.Response(w, listResponse, http.StatusOK)
}

//View : http handler for retrieve user by id
func (u *Users) View(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		return fmt.Errorf("type casting: %v", err)
	}

	var user models.User
	err = user.Get(u.Db, int64(id))
	if err != nil {
		return fmt.Errorf("Get user: %v", err)
	}

	var response response.UserResponse
	response.Transform(&user)
	return api.Response(w, response, http.StatusOK)
}

//Create : http handler for create new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) error {
	var userRequest request.NewUserRequest
	err := api.Decode(r, &userRequest)
	if err != nil {
		return fmt.Errorf("decode user: %v", err)
	}

	if userRequest.Password != userRequest.RePassword {
		return errors.New("Password not match")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Generate password: %v", err)
	}

	userRequest.Password = string(pass)
	user := userRequest.Transform()
	err = user.Create(u.Db)
	if err != nil {
		return fmt.Errorf("Create user: %v", err)
	}

	var response response.UserResponse
	response.Transform(user)
	return api.Response(w, response, http.StatusCreated)
}

//Update : http handler for update user by id
func (u *Users) Update(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		return fmt.Errorf("Type casting paramID: %v", err)
	}

	var user models.User
	err = user.Get(u.Db, int64(id))
	if err != nil {
		return fmt.Errorf("Get user: %v", err)
	}

	var userRequest request.UserRequest
	err = api.Decode(r, &userRequest)
	if err != nil {
		return fmt.Errorf("Decode user: %v", err)
	}

	if len(userRequest.Password) > 0 && userRequest.Password != userRequest.RePassword {
		return errors.New("Password not match")
	}

	if len(userRequest.Password) > 0 {
		pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("Generate password: %v", err)
		}

		userRequest.Password = string(pass)
	}

	if userRequest.ID <= 0 {
		userRequest.ID = user.ID
	}
	userUpdate := userRequest.Transform(&user)
	err = userUpdate.Update(u.Db)
	if err != nil {
		return fmt.Errorf("Update user: %v", err)
	}

	var response response.UserResponse
	response.Transform(userUpdate)
	return api.Response(w, response, http.StatusOK)
}

//Delete : http handler for delete user by id
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		return fmt.Errorf("Type casting paramID: %v", err)
	}

	var user models.User
	err = user.Get(u.Db, int64(id))
	if err != nil {
		return fmt.Errorf("Get user: %v", err)
	}

	err = user.Delete(u.Db)
	if err != nil {
		return fmt.Errorf("Delete user: %v", err)
	}

	return api.Response(w, nil, http.StatusNoContent)
}
