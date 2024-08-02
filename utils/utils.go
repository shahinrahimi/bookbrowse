package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// WrireJSON write data as application/json to responseWriter
func WriteJSON(rw http.ResponseWriter, status int, data any) error {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	return json.NewEncoder(rw).Encode(data)
}

// ToJSON serializes the given interface into a string based JSON format
func ToJSON(i interface{}, w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}

// FromJSON deserializes the object from JSON string
func FromJSON(i interface{}, r io.Reader) error {
	return json.NewDecoder(r).Decode(i)
}

// ParseQueryParamsPage parse request and return page number
// if error or not any params defined it will return default 1
func ParseQueryParamsPage(r *http.Request) int {
	if p := r.URL.Query().Get("page"); p != "" {
		parsed, err := strconv.Atoi(p)
		if err == nil && parsed > 0 {
			return parsed
		}
	}
	// return default value
	return 1
}

// ParseQueryParamsLimit parse request and return limit number
// if error or not any params defined it will return default 10
// the max value is 100
func ParseQueryParamsLimit(r *http.Request) int {
	if p := r.URL.Query().Get("limit"); p != "" {
		parsed, err := strconv.Atoi(p)
		if err == nil && parsed > 0 {
			if parsed > 100 {
				return 100
			}
			return parsed
		}
	}
	// return default value
	return 10
}

// PrseIDVars return non zero negetive value
func ParseIDVars(r *http.Request) int {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id < 1 {
		return -1
	}
	return id
}
