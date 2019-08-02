package routing

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/controllers"
	"github.com/jacky-htg/go-services/services/api"
	"github.com/jmoiron/sqlx"
)

//API : hanlder api
func API(db *sqlx.DB, log *log.Logger) http.Handler {
	app := newApp(log)
	u := controllers.Users{Db: db, Log: log}

	app.Handle(http.MethodGet, "/users", u.List)
	app.Handle(http.MethodGet, "/users/{id}", u.View)
	app.Handle(http.MethodPost, "/users", u.Create)
	app.Handle(http.MethodPut, "/users/{id}", u.Update)
	app.Handle(http.MethodDelete, "/users/{id}", u.Delete)

	return app
}

func newApp(log *log.Logger) *App {
	return &App{
		log: log,
		mux: chi.NewRouter(),
	}
}

// Handler is the signature used by all application handlers in this service.
type Handler func(http.ResponseWriter, *http.Request) error

// App is the entrypoint into our application and what controls the context of
// each request. Feel free to add any configuration data/logic on this type.
type App struct {
	log *log.Logger
	mux *chi.Mux
}

// Handle associates a handler function with an HTTP Method and URL pattern.
func (a *App) Handle(method, url string, h Handler) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Call the handler and catch any propagated error.
		err := h(w, r)

		if err != nil {
			// Log the error.
			a.log.Printf("ERROR : %+v", err)

			// Response to the error.
			if err := api.ResponseError(w, err); err != nil {
				a.log.Printf("ERROR : %v", err)
			}
		}
	}
	a.mux.MethodFunc(method, url, fn)
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
