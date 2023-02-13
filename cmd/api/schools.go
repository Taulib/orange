// FileName cmd/api/ schools.go

package main

import (
	"fmt"
	"net/http"
)

func (app *application) createSchoolHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Created a School...")
}

func (app *application) showSchoolHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParams(r)
	if err != nil {
		http.NotFound(r)
		return
	}
	fmt.Fprintf(w, "show details of schools %d\n", id)

}
