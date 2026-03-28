package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Student model
type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

// GET
func (app *application) getStudents(w http.ResponseWriter, r *http.Request) {

	rows, err := app.db.Query("SELECT id, name, programme, year FROM students")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var s Student
		rows.Scan(&s.ID, &s.Name, &s.Programme, &s.Year)
		students = append(students, s)
	}

	json.NewEncoder(w).Encode(students)
}

// POST
func (app *application) createStudent(w http.ResponseWriter, r *http.Request) {

	var s Student

	json.NewDecoder(r.Body).Decode(&s)

	if err := validateStudent(s); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `INSERT INTO students (name, programme, year)
	          VALUES ($1,$2,$3) RETURNING id`

	app.db.QueryRow(query, s.Name, s.Programme, s.Year).Scan(&s.ID)

	json.NewEncoder(w).Encode(s)
}

// PUT
func (app *application) updateStudent(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var s Student
	json.NewDecoder(r.Body).Decode(&s)

	query := `
	UPDATE students
	SET name=$1, programme=$2, year=$3
	WHERE id=$4
	`

	_, err := app.db.Exec(query, s.Name, s.Programme, s.Year, id)
	if err != nil {
		http.Error(w, "Update failed", 500)
		return
	}

	w.Write([]byte("Updated"))
}

// DELETE
func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	_, err := app.db.Exec("DELETE FROM students WHERE id=$1", id)
	if err != nil {
		http.Error(w, "Delete failed", 500)
		return
	}

	w.Write([]byte("Deleted"))
}