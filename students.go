package main

import (
	"encoding/json" // convert Go data → JSON
	"log"
	"net/http"
)

// 🧱 Student struct (represents table structure)
type Student struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Programme string `json:"programme"`
	Year      int    `json:"year"`
}

// 📥 Handler: Get all students
func getStudents(w http.ResponseWriter, r *http.Request) {

	// SQL query
	rows, err := DB.Query("SELECT id, name, programme, year FROM students")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// slice to store results
	var students []Student

	// loop through rows
	for rows.Next() {
		var s Student

		err := rows.Scan(&s.ID, &s.Name, &s.Programme, &s.Year)
		if err != nil {
			log.Println(err)
			continue
		}

		students = append(students, s)
	}

	// return JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(students)
}