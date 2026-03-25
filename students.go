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
func getStudents(w http.ResponseWriter, r *http.Request) {

	rows, err := DB.Query("SELECT id, name, programme, year FROM students")
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