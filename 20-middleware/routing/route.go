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
	u := controllers.Users{Db: db, Log: log}

	app.Handle(http.MethodGet, "/users", u.List)
	app.Handle(http.MethodGet, "/users/{id}", u.View)
	app.Handle(http.MethodPost, "/users", u.Create)
	app.Handle(http.MethodPut, "/users/{id}", u.Update)
	app.Handle(http.MethodDelete, "/users/{id}", u.Delete)

	return app
}
