package middleware

import (
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
)

func Errors(log *log.Logger) api.Middleware {
	fn := func(before api.Handler) api.Handler {
		h := func(w http.ResponseWriter, r *http.Request) error {

			// Call the handler and catch any propagated error.
			err := before(w, r)

			if err != nil {
				// Log the error.
				log.Printf("ERROR : %+v", err)

				// Response to the error.
				if err := api.ResponseError(w, err); err != nil {
					return err
				}
			}

			// Return nil to indicate the error has been handled.
			return nil
		}

		return h
	}

	return fn
}
