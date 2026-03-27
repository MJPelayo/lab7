package main

import (
	"database/sql"
	"net/http"
)

// application struct
type application struct {
	db *sql.DB
}

func (app *application) routes(mux *http.ServeMux) {

	mux.HandleFunc("/", app.home)

	// students
	mux.HandleFunc("/students", app.studentsHandler)

	// courses
	mux.HandleFunc("/courses", app.coursesHandler)
}

// 🎯 students router
func (app *application) studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getStudents(w, r)
	case http.MethodPost:
		app.createStudent(w, r)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

// 🎯 courses router
func (app *application) coursesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getCourses(w, r)
	case http.MethodPost:
		app.createCourse(w, r)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

// home
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Lab 7 API 🚀"))
}