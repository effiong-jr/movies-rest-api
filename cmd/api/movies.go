package main

import (
	"net/http"

	"fmt"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Create a movie")

}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of movie %d\n", id)

}
