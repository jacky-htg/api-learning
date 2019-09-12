package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jmoiron/sqlx"
)

func Errors(db *sqlx.DB, log *log.Logger) api.Middleware {
	fn := func(before api.Handler) api.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			err := before(ctx, w, r)

			if err != nil {
				// Log the error.
				log.Printf("ERROR : %+v", err)

				// Response to the error.
				if err := api.ResponseError(ctx, w, err); err != nil {
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
