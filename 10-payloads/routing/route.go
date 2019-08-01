package routing

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/controllers"
	"github.com/jmoiron/sqlx"
)

//API : hanlder api
func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := newApp(log)
	u := controllers.Users{Db: db, Log: log}

	app.Handle(http.MethodGet, "/users", u.List)
	app.Handle(http.MethodGet, "/users/{id}", u.View)

	return app
}

func newApp(log *log.Logger) *App {
	return &App{
		log: log,
		mux: chi.NewRouter(),
	}
}

// App is the entrypoint into our application and what controls the context of
// each request. Feel free to add any configuration data/logic on this type.
type App struct {
	log *log.Logger
	mux *chi.Mux
}

// Handle associates a handler function with an HTTP Method and URL pattern.
func (a *App) Handle(method, url string, h http.HandlerFunc) {
	a.mux.MethodFunc(method, url, h)
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
