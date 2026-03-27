package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// 🧱 Student model
type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

// 📥 GET /students
func (app *application) getStudents(w http.ResponseWriter, r *http.Request) {

	rows, err := app.db.Query("SELECT id, name, programme, year FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var s Student

		err := rows.Scan(&s.ID, &s.Name, &s.Programme, &s.Year)
		if err != nil {
			log.Println(err)
			continue
		}

		students = append(students, s)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}

// 📤 POST /students
func (app *application) createStudent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var s Student

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Basic validation
	if s.Name == "" || s.Programme == "" || s.Year == 0 {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	// Insert into DB
	query := `
		INSERT INTO students (name, programme, year)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err = app.db.QueryRow(query, s.Name, s.Programme, s.Year).Scan(&s.ID)
	if err != nil {
		http.Error(w, "Failed to insert student", http.StatusInternalServerError)
		return
	}

	// Return created student
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}
