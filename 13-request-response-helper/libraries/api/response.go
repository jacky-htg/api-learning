package api

import (
	"encoding/json"
	"net/http"
)

// Response converts a Go value to JSON and sends it to the client.
func Response(w http.ResponseWriter, data interface{}, statusCode int) error {

	// Convert the response value to JSON.
	res, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Respond with the provided JSON.
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}
