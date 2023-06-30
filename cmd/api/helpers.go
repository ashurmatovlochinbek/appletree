package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
)

// Define a new type named envelope
type envelope map[string]interface{}


func (app *application) readIDParam(r *http.Request) (int64, error) {
	// Use the "ParamsFromContext() function to get the request context as a sclice"
	params := httprouter.ParamsFromContext(r.Context())
	//Get the value of the "id" parameter 
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id paramater")
	}
	return id, nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

