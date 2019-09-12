package controllers

import (
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/libraries/token"
	"github.com/jacky-htg/go-services/models"
	"github.com/jacky-htg/go-services/payloads/request"
	"github.com/jacky-htg/go-services/payloads/response"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

//Auths : struct for set Auths Dependency Injection
type Auths struct {
	Db  *sqlx.DB
	Log *log.Logger
}

//Login : http handler for login
func (u *Auths) Login(w http.ResponseWriter, r *http.Request) error {
	var loginRequest request.LoginRequest
	err := api.Decode(r, &loginRequest)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "decode auth")
	}

	uLogin := models.User{Username: loginRequest.Username}
	err = uLogin.GetByUsername(r.Context(), u.Db)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "call login")
	}

	err = bcrypt.CompareHashAndPassword([]byte(uLogin.Password), []byte(loginRequest.Password))
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "compare password")
	}

	token, err := token.ClaimToken(uLogin.Username)
	if err != nil {
		u.Log.Printf("ERROR : %+v", err)
		return errors.Wrap(err, "claim token")
	}

	var response response.TokenResponse
	response.Token = token

	return api.ResponseOK(r.Context(), w, response, http.StatusOK)
}
