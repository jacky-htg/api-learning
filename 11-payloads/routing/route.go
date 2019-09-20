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

	return app
}
