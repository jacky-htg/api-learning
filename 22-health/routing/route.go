package routing

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/middleware"
	auth "github.com/jacky-htg/go-services/packages/auth/controllers"
	health "github.com/jacky-htg/go-services/packages/profiling/controllers"
)

func mid(db *sql.DB, log *log.Logger) []api.Middleware {
	var mw []api.Middleware
	mw = append(mw, middleware.Auths(db, log, []string{"/login", "/health"}))

	return mw
}

//API : hanlder api
func API(db *sql.DB, log *log.Logger) http.Handler {
	app := api.NewApp(log, mid(db, log)...)

	// Health Routing
	{
		check := health.Checks{Db: db}
		app.Handle(http.MethodGet, "/health", check.Health)
	}

	// Auth Routing
	{
		auth := auth.Auths{Db: db, Log: log}
		app.Handle(http.MethodPost, "/login", auth.Login)
	}

	// Users Routing
	{
		user := auth.Users{Db: db, Log: log}
		app.Handle(http.MethodGet, "/users", user.List)
		app.Handle(http.MethodGet, "/users/:id", user.View)
		app.Handle(http.MethodPost, "/users", user.Create)
		app.Handle(http.MethodPut, "/users/:id", user.Update)
		app.Handle(http.MethodDelete, "/users/:id", user.Delete)
	}

	// Roles Routing
	{
		roles := auth.Roles{Db: db, Log: log}
		app.Handle(http.MethodGet, "/roles", roles.List)
		app.Handle(http.MethodGet, "/roles/:id", roles.View)
		app.Handle(http.MethodPost, "/roles", roles.Create)
		app.Handle(http.MethodPut, "/roles/:id", roles.Update)
		app.Handle(http.MethodDelete, "/roles/:id", roles.Delete)
		app.Handle(http.MethodPost, "/roles/:id/access/:access_id", roles.Grant)
		app.Handle(http.MethodDelete, "/roles/:id/access/:access_id", roles.Revoke)
	}

	// Access Routing
	{
		access := auth.Access{Db: db, Log: log}
		app.Handle(http.MethodGet, "/access", access.List)
	}

	return app
}
