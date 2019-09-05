package routing

import (
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/controllers"
	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/middleware"
	"github.com/jmoiron/sqlx"
)

//API : handler api
func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := api.NewApp(log, middleware.Errors(log))

	// Auth Routing
	auth := controllers.Auths{Db: db, Log: log}
	app.Handle(http.MethodPost, "/login", auth.Login)

	// Users Routing
	u := controllers.Users{Db: db, Log: log}
	app.Handle(http.MethodGet, "/users", u.List)
	app.Handle(http.MethodGet, "/users/{id}", u.View)
	app.Handle(http.MethodPost, "/users", u.Create)
	app.Handle(http.MethodPut, "/users/{id}", u.Update)
	app.Handle(http.MethodDelete, "/users/{id}", u.Delete)

	// Roles Routing
	roles := controllers.Roles{Db: db, Log: log}
	app.Handle(http.MethodGet, "/roles", roles.List)
	app.Handle(http.MethodGet, "/roles/{id}", roles.View)
	app.Handle(http.MethodPost, "/roles", roles.Create)
	app.Handle(http.MethodPut, "/roles/{id}", roles.Update)
	app.Handle(http.MethodDelete, "/roles/{id}", roles.Delete)
	app.Handle(http.MethodPost, "/roles/{id}/access/{access_id}", roles.Grant)
	app.Handle(http.MethodDelete, "/roles/{id}/access/{access_id}", roles.Revoke)

	// Access Routing
	access := controllers.Access{Db: db, Log: log}
	app.Handle(http.MethodGet, "/access", access.List)

	return app
}
