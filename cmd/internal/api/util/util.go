package util

import (
	"encoding/json"
	"net/http"
)

// WriteToJSON writes the value as JSON to the response writer
func WriteToJSON(w http.ResponseWriter, s int, v any) error {

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Set the status code for the request
	w.WriteHeader(s)

	// Encode the value v as JSON and write it to the response
	return json.NewEncoder(w).Encode(v)
}
