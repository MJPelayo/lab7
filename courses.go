package main

import (
	"encoding/json"
	"net/http"
)

type Course struct {
	ID         int    `json:"id"`
	Code       string `json:"code"`
	Title      string `json:"title"`
	Department string `json:"department"`
	Instructor string `json:"instructor"`
	Credits    int    `json:"credits"`
	Capacity   int    `json:"capacity"`
	Enrolled   int    `json:"enrolled"`
}

func (app *application) getCourses(w http.ResponseWriter, r *http.Request) {

	rows, _ := app.db.Query("SELECT id, code, title, department, instructor, credits, capacity, enrolled FROM courses")
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var c Course
		rows.Scan(&c.ID, &c.Code, &c.Title, &c.Department, &c.Instructor, &c.Credits, &c.Capacity, &c.Enrolled)
		courses = append(courses, c)
	}

	json.NewEncoder(w).Encode(courses)
}

func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {

	var c Course
	json.NewDecoder(r.Body).Decode(&c)

	if err := validateCourse(c); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `INSERT INTO courses (code, title, department, instructor, credits, capacity, enrolled)
			  VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`

	app.db.QueryRow(query, c.Code, c.Title, c.Department, c.Instructor, c.Credits, c.Capacity, c.Enrolled).Scan(&c.ID)

	json.NewEncoder(w).Encode(c)
}
