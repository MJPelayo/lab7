package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)
// Student struct represents a student in the system
type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

// GET ALL STUDENTS (VIEW)

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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		students = append(students, s)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(students)
}
// CREATE STUDENT (ADD)


func (app *application) createStudent(w http.ResponseWriter, r *http.Request) {
	var s Student
	
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := validateStudent(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `INSERT INTO students (name, programme, year) VALUES ($1,$2,$3) RETURNING id`
	err = app.db.QueryRow(query, s.Name, s.Programme, s.Year).Scan(&s.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}
// UPDATE STUDENT

func (app *application) updateStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var s Student
	err = json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	_, err = app.db.Exec("UPDATE students SET name=$1, programme=$2, year=$3 WHERE id=$4",
		s.Name, s.Programme, s.Year, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Student updated"))
}

// DELETE STUDENT


func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = app.db.Exec("DELETE FROM students WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Student deleted"))
}
