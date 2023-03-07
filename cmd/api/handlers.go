package main

import (
	"log"
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

	_ = app.writeJSON(w, http.StatusOK, payload)
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
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	//Read json payload

	//validate user against database

	//check password

	//create a jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	//generate tokens
	tokens, err := app.auth.GenerateTokenPair(&u)
	log.Println(err)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	log.Println(tokens.Token)

	w.Write([]byte(tokens.Token))
}
