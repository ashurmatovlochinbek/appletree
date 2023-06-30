package main

import (
	"appletree/internal/data"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// createSchoolHandler for the "POST /v1/schools endpoint"
func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string 		`json:"name"`
		Level string 		`json:"level"`
		Contact string 		`json:"contact"`
		Phone string		`json:"phone"`
		Email string		`json:"email"`
		Website string		`json:"website"`
		Address string		`json:"address"`
		Mode []string		`json:"mode"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err)
		return
	}
	
	fmt.Fprintf(w, "%+v\n,", input)
}

// createSchoolHandler for the "GET /v1/schools/:id endpoint"
func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
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

	err = app.writeJSON(w, http.StatusOK, envelope{"school":school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

