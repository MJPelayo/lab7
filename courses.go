package main

import (
	"encoding/json"
	"net/http"
)

// 🧱 Course model
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

// 📥 GET /courses
func (app *application) getCourses(w http.ResponseWriter, r *http.Request) {

	rows, err := app.db.Query("SELECT id, code, title, department, instructor, credits, capacity, enrolled FROM courses")
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var courses []Course

	for rows.Next() {
		var c Course

		err := rows.Scan(
			&c.ID,
			&c.Code,
			&c.Title,
			&c.Department,
			&c.Instructor,
			&c.Credits,
			&c.Capacity,
			&c.Enrolled,
		)
		if err != nil {
			continue
		}

		courses = append(courses, c)
	}

	json.NewEncoder(w).Encode(courses)
}

// 📤 POST /courses
func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", 405)
		return
	}

	var c Course

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	// validation
	if err := validateCourse(c); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	query := `
	INSERT INTO courses (code, title, department, instructor, credits, capacity, enrolled)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id
	`

	err = app.db.QueryRow(
		query,
		c.Code,
		c.Title,
		c.Department,
		c.Instructor,
		c.Credits,
		c.Capacity,
		c.Enrolled,
	).Scan(&c.ID)

	if err != nil {
		http.Error(w, "Insert failed", 500)
		return
	}

	json.NewEncoder(w).Encode(c)
}
