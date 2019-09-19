package api

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// App is the entrypoint into our application and what controls the context of
// each request. Feel free to add any configuration data/logic on this type.
type App struct {
	log *log.Logger
	mux *httprouter.Router
}

// Handle associates a httprouter Handle function with an HTTP Method and URL pattern.
func (a *App) Handle(method, url string, h httprouter.Handle) {
	a.mux.Handle(method, url, h)
}

// ServeHTTP implements the http.Handler interface.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

//NewApp is function to create new App
func NewApp(log *log.Logger) *App {
	return &App{
		log: log,
		mux: httprouter.New(),
	}
}
