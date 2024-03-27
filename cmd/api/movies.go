package main

import (
	"net/http"

	"fmt"
	"time"

	"github.com/effiong-jr/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Create a movie")

}

func (app *application) getMovieHandler(w http.ResponseWriter, r *http.Request) {

	movie := data.Movie{
		ID:        1,
		Title:     "Inception",
		Year:      2010,
		Runtime:   148,
		Genres:    []string{"Sci-fi", "Action"},
		CreatedAt: time.Now(),
		Version:   1,
	}

	err := app.writeJSON(w, http.StatusOK, movie, nil)

	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered and error and could not process your request", http.StatusInternalServerError)
		return
	}

}
