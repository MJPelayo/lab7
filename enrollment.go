package main

import (
	"net/http"
	"strconv"
)

// 🎯 POST /enroll?student_id=1&course_id=2
func (app *application) enrollStudent(w http.ResponseWriter, r *http.Request) {

	sid, _ := strconv.Atoi(r.URL.Query().Get("student_id"))
	cid, _ := strconv.Atoi(r.URL.Query().Get("course_id"))

	query := `
	INSERT INTO enrollments (student_id, course_id)
	VALUES ($1,$2)
	`

	_, err := app.db.Exec(query, sid, cid)
	if err != nil {
		http.Error(w, "Enrollment failed", 500)
		return
	}

	w.Write([]byte("Student enrolled"))
}