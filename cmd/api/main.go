package main

import (
	"backend/internal/repository"
	"backend/internal/repository/dbrepo"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	DSN    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	//set application config
	var app application

	//read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgresdb password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	fmt.Println("dSN is: ", app.DSN)

	//connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	// Connecting the DatabaseRepo interface to the PostgresDBRepo struct
	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	//start a webserver
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatal(err)
	}
}
