package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

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
