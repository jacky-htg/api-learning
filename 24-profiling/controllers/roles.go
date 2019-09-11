package controllers

import (
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
)

//Roles : struct for set Roles Dependency Injection
type Roles struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//List : http handler for returning list of roles
func (u *Roles) List(w http.ResponseWriter, r *http.Request) error {
	var role models.Role
	list, err := role.List(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "getting roles list")
	}

	var listResponse []*response.RoleResponse
	for _, role := range list {
		var roleResponse response.RoleResponse
		roleResponse.Transform(&role)
		listResponse = append(listResponse, &roleResponse)
	}

	return api.ResponseOK(w, listResponse, http.StatusOK)
}

//View : http handler for retrieve role by id
func (u *Roles) View(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting")
	}

	var role models.Role
	role.ID = uint32(id)
	err = role.Get(r.Context(), u.Db)

	if err == sql.ErrNoRows {
		u.Log.Printf("ERROR : %+v", err)
		return api.NotFoundError(err, "")
	}

	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get Role")
	}

	var response response.RoleResponse
	response.Transform(&role)
	return api.ResponseOK(w, response, http.StatusOK)
}

//Create : http handler for create new role
func (u *Roles) Create(w http.ResponseWriter, r *http.Request) error {
	var roleRequest request.NewRoleRequest
	err := api.Decode(r, &roleRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode role")
	}

	role := roleRequest.Transform()
	err = role.Create(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Create Role")
	}

	var response response.RoleResponse
	response.Transform(role)
	return api.ResponseOK(w, response, http.StatusCreated)
}

//Update : http handler for update role by id
func (u *Roles) Update(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var role models.Role
	role.ID = uint32(id)
	err = role.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get Role")
	}

	var roleRequest request.RoleRequest
	err = api.Decode(r, &roleRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Decode Role")
	}

	if roleRequest.ID <= 0 {
		roleRequest.ID = role.ID
	}
	roleUpdate := roleRequest.Transform(&role)
	err = roleUpdate.Update(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Update role")
	}

	var response response.RoleResponse
	response.Transform(roleUpdate)
	return api.ResponseOK(w, response, http.StatusOK)
}

//Delete : http handler for delete role by id
func (u *Roles) Delete(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	var role models.Role
	role.ID = uint32(id)
	err = role.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get role")
	}

	isDelete, err := role.Delete(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Delete role")
	}

	if !isDelete {
		err = errors.New("error delete role")
		u.Log.Printf("ERROR : %+v", err)
		return err
	}

	return api.ResponseOK(w, nil, http.StatusNoContent)
}

//Grant : http handler for grant access to role
func (u *Roles) Grant(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")
	paramAccessID := chi.URLParam(r, "access_id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	accessID, err := strconv.Atoi(paramAccessID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramAccessID")
	}

	var role models.Role
	role.ID = uint32(id)
	err = role.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get role")
	}

	var access models.Access
	access.ID = uint32(accessID)
	tx := u.Db.MustBegin()
	err = access.Get(r.Context(), tx)
	if err != nil {
		tx.Rollback()
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get access")
	}
	tx.Commit()

	err = role.Grant(r.Context(), u.Db, access.ID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Grant role")
	}

	return api.ResponseOK(w, nil, http.StatusOK)
}

//Revoke : http handler for revoke access from role
func (u *Roles) Revoke(w http.ResponseWriter, r *http.Request) error {
	paramID := chi.URLParam(r, "id")
	paramAccessID := chi.URLParam(r, "access_id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramID")
	}

	accessID, err := strconv.Atoi(paramAccessID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "type casting paramAccessID")
	}

	var role models.Role
	role.ID = uint32(id)
	err = role.Get(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get role")
	}

	var access models.Access
	access.ID = uint32(accessID)
	tx := u.Db.MustBegin()
	err = access.Get(r.Context(), tx)
	if err != nil {
		tx.Rollback()
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Get access")
	}
	tx.Commit()

	err = role.Revoke(r.Context(), u.Db, access.ID)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "Revoke role")
	}

	return api.ResponseOK(w, nil, http.StatusOK)
}
