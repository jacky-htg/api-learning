package routing

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/packages/auth/controllers"
)

//API : hanlder api
func API(db *sql.DB, log *log.Logger) http.Handler {
	app := api.NewApp(log)
	user := controllers.Users{Db: db, Log: log}

	app.Handle(http.MethodGet, "/users", user.List)
	app.Handle(http.MethodGet, "/users/:id", user.View)
	app.Handle(http.MethodPost, "/users", user.Create)
	app.Handle(http.MethodPut, "/users/:id", user.Update)
	app.Handle(http.MethodDelete, "/users/:id", user.Delete)

	return app
}
