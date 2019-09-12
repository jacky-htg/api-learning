package middleware

import (
	"bytes"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/jacky-htg/go-services/libraries/api"
)

// Logger writes some information about the request to the logs in the
// format: (200) GET /foo -> (request) -> (response) -> IP ADDR (latency)
func Logger(log *log.Logger) api.Middleware {

	// This is the actual middleware function to be executed.
	f := func(before api.Handler) api.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			v, ok := ctx.Value(api.KeyValues).(*api.Values)
			if !ok {
				return errors.New("web value missing from context")
			}

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				return err
			}

			r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			err = before(ctx, w, r)

			// You can save it in your nosql database
			log.Printf("(%d) : %s %s -> %s -> %s -> %s (%s)",
				v.StatusCode,
				r.Method, r.URL.Path,
				body,
				v.Response,
				r.RemoteAddr, time.Since(v.Start),
			)

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}

	return f
}
