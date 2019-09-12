package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/jacky-htg/go-services/libraries/api"
	"github.com/jacky-htg/go-services/libraries/array"
	"github.com/jacky-htg/go-services/models"
	"github.com/jmoiron/sqlx"
)

func Auths(db *sqlx.DB, log *log.Logger, allow []string) api.Middleware {
	fn := func(before api.Handler) api.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			rctx := chi.RouteContext(ctx)
			url := r.URL.String()
			controller := strings.Split(url, "/")[1]
			var access models.Access
			isAuth, err := access.IsAuth(ctx, db, r.Header.Get("Token"), controller, rctx.RouteMethod+" "+rctx.RoutePatterns[0])

			var astr array.ArrStr
			inArray, _ := astr.InArray(url, allow)
			if !inArray && (err != nil || !isAuth) {
				err = api.ForbiddenError(errors.New("Forbidden auth"), "")
			} else {
				err = before(ctx, w, r)
			}

			return err
		}

		return h
	}

	return fn
}
