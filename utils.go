package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	maxBytes := int64(1048576)

	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain one single json value")
	}

	return nil

}

// WriteJSON a helper function that allows us to write payload to JSON format
func WriteJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}, headers ...http.Header) error {

	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(out)
	return nil
}
