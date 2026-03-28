package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Course struct represents a course in the system
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


// GET ALL COURSES (VIEW)


func (app *application) getCourses(w http.ResponseWriter, r *http.Request) {

	// Query all courses from database
	rows, err := app.db.Query(`
		SELECT id, code, title, department, instructor, credits, capacity, enrolled 
		FROM courses`)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	var courses []Course

	// Loop through result set
	for rows.Next() {
		var c Course
		err := rows.Scan(&c.ID, &c.Code, &c.Title, &c.Department, &c.Instructor, &c.Credits, &c.Capacity, &c.Enrolled)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		courses = append(courses, c)
	}

	// Return JSON response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}


// CREATE COURSE (ADD)


func (app *application) createCourse(w http.ResponseWriter, r *http.Request) {

	var c Course

	// Decode incoming JSON
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	// Validate input
	if err := validateCourse(c); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Insert into DB
	query := `
	INSERT INTO courses 
	(code, title, department, instructor, credits, capacity, enrolled)
	VALUES ($1,$2,$3,$4,$5,$6,$7)
	RETURNING id`

	err = app.db.QueryRow(query,
		c.Code, c.Title, c.Department, c.Instructor,
		c.Credits, c.Capacity, c.Enrolled).Scan(&c.ID)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Return created course
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}


//UPDATE COURSE


func (app *application) updateCourse(w http.ResponseWriter, r *http.Request) {

	// Get course ID from query param
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	var c Course

	// Decode updated data
	err = json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid JSON", 400)
		return
	}

	// Update DB record
	_, err = app.db.Exec(`
		UPDATE courses 
		SET code=$1, title=$2, department=$3, instructor=$4, credits=$5, capacity=$6, enrolled=$7
		WHERE id=$8`,
		c.Code, c.Title, c.Department, c.Instructor,
		c.Credits, c.Capacity, c.Enrolled, id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Course updated"))
}

// DELETE COURSE


func (app *application) deleteCourse(w http.ResponseWriter, r *http.Request) {

	// Get course ID
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	// Delete from DB
	_, err = app.db.Exec("DELETE FROM courses WHERE id=$1", id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Course deleted"))
}