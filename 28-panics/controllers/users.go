package controllers

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/models"
	"github.com/jacky-htg/go-services/payloads/request"
	"github.com/jacky-htg/go-services/payloads/response"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//Users : struct for set Users Dependency Injection
type Users struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of users
func (u *Users) List(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var user models.User
	list, err := user.List(ctx, u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "getting users list")
	}

	var listResponse []*response.UserResponse
	for _, user := range list {
		var userResponse response.UserResponse
		userResponse.Transform(&user)
		listResponse = append(listResponse, &userResponse)
	}

	return api.ResponseOK(ctx, w, listResponse, http.StatusOK)
}

//View : http handler for retrieve user by id
func (u *Users) View(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting")
	}

	var user models.User
	user.ID = uint64(id)
	err = user.Get(ctx, u.Db)

	if err == sql.ErrNoRows {
		u.Log.Printf("ERROR : %+v", err)
		return api.NotFoundError(err, "")
	}

	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get User")
	}

	var response response.UserResponse
	response.Transform(&user)
	return api.ResponseOK(ctx, w, response, http.StatusOK)
}

//Create : http handler for create new user
func (u *Users) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var userRequest request.NewUserRequest
	err := api.Decode(r, &userRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode user")
	}

	if userRequest.Password != userRequest.RePassword {
		err = errors.New("Password not match")
		u.Log.Printf("ERROR : %+v", err)
		return api.BadRequestError(err, "Password not match")
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Generate password")
	}

	userRequest.Password = string(pass)
	user := userRequest.Transform()
	tx := u.Db.MustBegin()
	err = user.Create(ctx, tx)
	if err != nil {
		tx.Rollback()
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Create User")
	}

	tx.Commit()

	var response response.UserResponse
	response.Transform(user)
	return api.ResponseOK(ctx, w, response, http.StatusCreated)
}

//Update : http handler for update user by id
func (u *Users) Update(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var user models.User
	user.ID = uint64(id)
	err = user.Get(ctx, u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get User")
	}

	var userRequest request.UserRequest
	err = api.Decode(r, &userRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Decode User")
	}

	if len(userRequest.Password) > 0 && userRequest.Password != userRequest.RePassword {
		err = errors.New("Password not match")
		u.Log.Printf("ERROR : %+v", err)
		return api.BadRequestError(err, "Password not match")
	}

	if len(userRequest.Password) > 0 {
		pass, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
		if err != nil {
			u.Log.Printf("ERROR : %+v", err)
			return errors.Wrap(err, "Generate password")
		}

		userRequest.Password = string(pass)
	}

	if userRequest.ID <= 0 {
		userRequest.ID = user.ID
	}
	userUpdate := userRequest.Transform(&user)
	tx := u.Db.MustBegin()
	err = userUpdate.Update(ctx, tx)
	if err != nil {
		tx.Rollback()
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Update user")
	}

	tx.Commit()

	var response response.UserResponse
	response.Transform(userUpdate)
	return api.ResponseOK(ctx, w, response, http.StatusOK)
}

//Delete : http handler for delete user by id
func (u *Users) Delete(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var user models.User
	user.ID = uint64(id)
	err = user.Get(ctx, u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get user")
	}

	isDelete, err := user.Delete(ctx, u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Delete user")
	}

	if !isDelete {
		err = errors.New("error delete user")
		u.Log.Printf("ERROR : %+v", err)
		return err
	}

	return api.ResponseOK(ctx, w, nil, http.StatusNoContent)
}
