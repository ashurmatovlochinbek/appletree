package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createSchoolHandler for the "POST /v1/schools endpoint"
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new school")

}

// createSchoolHandler for the "GET /v1/schools/:id endpoint"
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	// Use the "ParamsFromContext() function to get the request context as a sclice"
	params := httprouter.ParamsFromContext(r.Context())
	//Get the value of the "id" parameter 
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details for school %d\n", id)
}