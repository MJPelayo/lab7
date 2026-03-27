package main

import (
	"database/sql"
	"net/http"
)

// Git Commit: refactor routes using application struct

// application struct (holds dependencies)
type application struct {
	db *sql.DB
}

// register all routes
func (app *application) routes(mux *http.ServeMux) {

	// home route
	mux.HandleFunc("/", app.home)

	// student routes
	mux.HandleFunc("/students", app.getStudents)
	mux.HandleFunc("/students", app.createStudent)

	// later we will add:
	// PUT /students/{id}
	// DELETE /students/{id}
}

// home handler
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Lab 7 API 🚀"))
}