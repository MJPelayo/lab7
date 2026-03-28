package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

func (app *application) getStudents(w http.ResponseWriter, r *http.Request) {

	rows, _ := app.db.Query("SELECT id, name, programme, year FROM students")
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var s Student
		rows.Scan(&s.ID, &s.Name, &s.Programme, &s.Year)
		students = append(students, s)
	}

	json.NewEncoder(w).Encode(students)
}

func (app *application) createStudent(w http.ResponseWriter, r *http.Request) {

	var s Student
	json.NewDecoder(r.Body).Decode(&s)

	if err := validateStudent(s); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `INSERT INTO students (name, programme, year) VALUES ($1,$2,$3) RETURNING id`
	app.db.QueryRow(query, s.Name, s.Programme, s.Year).Scan(&s.ID)

	json.NewEncoder(w).Encode(s)
}

func (app *application) updateStudent(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	var s Student
	json.NewDecoder(r.Body).Decode(&s)

	app.db.Exec("UPDATE students SET name=$1, programme=$2, year=$3 WHERE id=$4",
		s.Name, s.Programme, s.Year, id)

	w.Write([]byte("Updated"))
}

func (app *application) deleteStudent(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	app.db.Exec("DELETE FROM students WHERE id=$1", id)

	w.Write([]byte("Deleted"))
}
