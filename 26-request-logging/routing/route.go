package routing

import (
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/controllers"
	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/middleware"
	"github.com/jmoiron/sqlx"
)

func mid(db *sqlx.DB, log *log.Logger) []api.Middleware {
	var mw []api.Middleware
	mw = append(mw, middleware.Errors(db, log))
	mw = append(mw, middleware.Auths(db, log, []string{"/login", "/health"}))
	mw = append(mw, middleware.Metrics())
	mw = append(mw, middleware.Logger(log))

	return mw
}

//API : handler api
func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := api.NewApp(db, log, mid(db, log)...)

	// Health Routing
	{
		check := controllers.Checks{Db: db}
		app.Handle(http.MethodGet, "/health", check.Health)
	}

	// Auth Routing
	{
		auth := controllers.Auths{Db: db, Log: log}
		app.Handle(http.MethodPost, "/login", auth.Login)
	}

	// Users Routing
	{
		users := controllers.Users{Db: db, Log: log}
		app.Handle(http.MethodGet, "/users", users.List)
		app.Handle(http.MethodGet, "/users/{id}", users.View)
		app.Handle(http.MethodPost, "/users", users.Create)
		app.Handle(http.MethodPut, "/users/{id}", users.Update)
		app.Handle(http.MethodDelete, "/users/{id}", users.Delete)
	}

	// Roles Routing
	{
		roles := controllers.Roles{Db: db, Log: log}
		app.Handle(http.MethodGet, "/roles", roles.List)
		app.Handle(http.MethodGet, "/roles/{id}", roles.View)
		app.Handle(http.MethodPost, "/roles", roles.Create)
		app.Handle(http.MethodPut, "/roles/{id}", roles.Update)
		app.Handle(http.MethodDelete, "/roles/{id}", roles.Delete)
		app.Handle(http.MethodPost, "/roles/{id}/access/{access_id}", roles.Grant)
		app.Handle(http.MethodDelete, "/roles/{id}/access/{access_id}", roles.Revoke)
	}

	// Access Routing
	{
		access := controllers.Access{Db: db, Log: log}
		app.Handle(http.MethodGet, "/access", access.List)
	}

	return app
}
