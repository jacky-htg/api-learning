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
	u := controllers.Users{Db: db, Log: log}

	app.Handle(http.MethodGet, "/users", u.List)
	app.Handle(http.MethodGet, "/users/:id", u.View)

	return app
}
