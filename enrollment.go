package main

import (
	"net/http"
	"strconv"
)

func (app *application) enrollStudent(w http.ResponseWriter, r *http.Request) {

	sid, _ := strconv.Atoi(r.URL.Query().Get("student_id"))
	cid, _ := strconv.Atoi(r.URL.Query().Get("course_id"))

	app.db.Exec("INSERT INTO enrollments (student_id, course_id) VALUES ($1,$2)", sid, cid)

	w.Write([]byte("Enrolled"))
}
