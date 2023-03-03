package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	// var movies []models.Movie

	// rd, _ := time.Parse("2006-01-02", "1986-03-07")

	// highlander := models.Movie{
	// 	ID:          1,
	// 	Title:       "Highlander",
	// 	ReleaseDate: rd,
	// 	MPAARating:  "R",
	// 	RunTime:     116,
	// 	Description: "A very nice movie",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }

	// movies = append(movies, highlander)

	// rd, _ = time.Parse("2006-01-02", "1999-07-08")

	// DDLJ := models.Movie{
	// 	ID:          2,
	// 	Title:       "Dilwale Dulhaniya le Jaenge",
	// 	ReleaseDate: rd,
	// 	MPAARating:  "PG-13",
	// 	RunTime:     140,
	// 	Description: "A very romantic movie",
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }

	// movies = append(movies, DDLJ)

	movies, err := app.DB.AllMovies()
	if err != nil {
		fmt.Println(err)
		return
	}

	out, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}