package main

import (
	"database/sql"
	"net/http"
)

type application struct {
	db *sql.DB
}

func (app *application) routes(mux *http.ServeMux) {

	mux.HandleFunc("/", app.home)

	mux.HandleFunc("/students", app.studentsHandler)
	mux.HandleFunc("/courses", app.coursesHandler)
	mux.HandleFunc("/enroll", app.enrollStudent)	
	
}

func (app *application) studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getStudents(w, r)
	case http.MethodPost:
		app.createStudent(w, r)
	case http.MethodPut:
		app.updateStudent(w, r)
	case http.MethodDelete:
		app.deleteStudent(w, r)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

func (app *application) coursesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getCourses(w, r)
	case http.MethodPost:
		app.createCourse(w, r)
	case http.MethodPut:
		app.updateCourse(w, r)
	case http.MethodDelete:
		app.deleteCourse(w, r)
	default:
		http.Error(w, "Method not allowed", 405)
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Lab 7 API 🚀"))
}
