package main

import (
	"appletree/internal/data"
	"fmt"
	"net/http"
	"time"
)

// createSchoolHandler for the "POST /v1/schools endpoint"
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new school")

}

// createSchoolHandler for the "GET /v1/schools/:id endpoint"
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	school := data.School{
		ID: id,
		CreatedAt: time.Now(),
		Name: "Apple Tree",
		Level: "High School",
		Contact: "Anna Smith",
		Phone: "601-4411",
		Address: "14 Apple street",
		Mode: []string{"blended", "online"},
		Version: 1,
	}

	err = app.writeJSON(w, http.StatusOK, school, nil)
	if err != nil {
		app.logger.Panicln(err)
		http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	


}

