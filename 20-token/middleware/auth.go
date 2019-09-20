package middleware

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/jacky-htg/go-services/libraries/api"
)

func Auths(db *sql.DB, log *log.Logger, allow []string) api.Middleware {
	fn := func(before api.Handler) api.Handler {
		h := func(w http.ResponseWriter, r *http.Request) {
			var isAuth bool

			// hardcode athorization for true.
			// upcoming chapter, this line will execute RBAC checking
			isAuth = true

			if !isAuth {
				api.ResponseError(w, api.ErrForbidden(errors.New("Forbidden"), ""))
			} else {
				before(w, r)
			}
		}

		return h
	}

	return fn
}
