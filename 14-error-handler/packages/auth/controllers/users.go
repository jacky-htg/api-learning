package controllers

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/packages/auth/controllers/request"
	"github.com/jacky-htg/go-services/packages/auth/controllers/response"
	"github.com/jacky-htg/go-services/packages/auth/models"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

//Users : struct for set Users Dependency Injection
type Users struct {
	Db  *sql.DB
	Log *log.Logger
}

//List : http handler for returning list of users
func (u *Users) List(w http.ResponseWriter, r *http.Request) {
	var user models.User
	list, err := user.List(u.Db)
	if err != nil {
		u.Log.Printf("error call list users: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var listResponse []*response.UserResponse
	for _, user := range list {
		var userResponse response.UserResponse
		userResponse.Transform(&user)
		listResponse = append(listResponse, &userResponse)
	}

	api.ResponseOK(w, listResponse, http.StatusOK)
}

//View : http handler for retrieve user by id
func (u *Users) View(w http.ResponseWriter, r *http.Request) {
	paramID := r.Context().Value("ps").(httprouter.Params).ByName("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var user models.User
	err = user.Get(u.Db, int64(id))
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var response response.UserResponse
	response.Transform(&user)
	api.ResponseOK(w, response, http.StatusOK)
}

//Create : http handler for create new user
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var userRequest request.NewUserRequest
	err := api.Decode(r, &userRequest)
	if err != nil {
		u.Log.Printf("error decode user: %s", err)
		api.ResponseError(w, err)
		return
	}

	if userRequest.Password != userRequest.RePassword {
		err = errors.New("Password not match")
		u.Log.Printf("error : %s", err)
		api.ResponseError(w, api.ErrBadRequest(err, ""))
		return
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("error generate password: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	userRequest.Password = string(pass)

	user := userRequest.Transform()
	err = user.Create(u.Db)
	if err != nil {
		u.Log.Printf("error call create user: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var response response.UserResponse
	response.Transform(user)
	api.ResponseOK(w, response, http.StatusCreated)
}

//Update : http handler for update user by id
func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	paramID := r.Context().Value("ps").(httprouter.Params).ByName("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var user models.User
	err = user.Get(u.Db, int64(id))
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var userRequest request.UserRequest
	err = api.Decode(r, &userRequest)
	if err != nil {
		u.Log.Printf("error decode user: %s", err)
		api.ResponseError(w, err)
		return
	}

	if len(userRequest.Password) > 0 && userRequest.Password != userRequest.RePassword {
		err = errors.New("Password not match")
		u.Log.Printf("error : %s", err)
		api.ResponseError(w, api.ErrBadRequest(err, ""))
		return
	}

	if len(userRequest.Password) > 0 {
		pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			u.Log.Printf("error generate password: %s", err)
			api.ResponseError(w, api.ErrInternal(err, ""))
			return
		}

		userRequest.Password = string(pass)
	}

	if userRequest.ID <= 0 {
		userRequest.ID = user.ID
	}
	userUpdate := userRequest.Transform(&user)
	err = userUpdate.Update(u.Db)
	if err != nil {
		u.Log.Printf("error call update user: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var response response.UserResponse
	response.Transform(userUpdate)
	api.ResponseOK(w, response, http.StatusOK)
}

//Delete : http handler for delete user by id
func (u *Users) Delete(w http.ResponseWriter, r *http.Request) {
	paramID := r.Context().Value("ps").(httprouter.Params).ByName("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("error type casting paramID: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	var user models.User
	err = user.Get(u.Db, int64(id))
	if err != nil {
		u.Log.Printf("error call list user: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	err = user.Delete(u.Db)
	if err != nil {
		u.Log.Printf("error call delete user: %s", err)
		api.ResponseError(w, api.ErrInternal(err, ""))
		return
	}

	api.ResponseOK(w, nil, http.StatusNoContent)
}
