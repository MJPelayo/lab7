package main

import (
	"log"
	"net/http"
)

// Git Commit: refactor application structure

func main() {

	// initialize database
	db := openDB()

	// create app struct (dependency injection)
	app := &application{
		db: db,
	}

	// create router
	mux := http.NewServeMux()

	// register routes
	app.routes(mux)

	log.Println("🚀 Server running on :4000")

	// start server
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
