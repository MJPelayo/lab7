package main

import (
	"encoding/json"
	"net/http"
)

// Enrollment struct represents relationship between student and course
type Enrollment struct {
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
}

// ENROLLMENTS HANDLER
func (app *application) enrollmentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		app.getEnrollments(w, r)
	case http.MethodPost:
		app.createEnrollment(w, r)
	case http.MethodDelete:
		app.deleteEnrollment(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// VIEW ALL ENROLLMENTS
func (app *application) getEnrollments(w http.ResponseWriter, r *http.Request) {
	// Query enrollments table
	rows, err := app.db.Query("SELECT student_id, course_id FROM enrollments")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var enrollments []Enrollment

	// Loop through results
	for rows.Next() {
		var e Enrollment
		err := rows.Scan(&e.StudentID, &e.CourseID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		enrollments = append(enrollments, e)
	}

	// Return JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enrollments)
}

// CREATE ENROLLMENT
func (app *application) createEnrollment(w http.ResponseWriter, r *http.Request) {
	var e Enrollment
	
	// Decode incoming JSON
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Insert into DB
	_, err = app.db.Exec(`
		INSERT INTO enrollments (student_id, course_id)
		VALUES ($1, $2)`,
		e.StudentID, e.CourseID)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Enrollment created"))
}

// DELETE ENROLLMENT
func (app *application) deleteEnrollment(w http.ResponseWriter, r *http.Request) {
	studentID := r.URL.Query().Get("student_id")
	courseID := r.URL.Query().Get("course_id")
	
	if studentID == "" || courseID == "" {
		http.Error(w, "Missing student_id or course_id", http.StatusBadRequest)
		return
	}

	_, err := app.db.Exec(`
		DELETE FROM enrollments
		WHERE student_id = $1 AND course_id = $2`,
		studentID, courseID)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Enrollment deleted"))
}