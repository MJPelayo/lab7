package main

import (
	"fmt"
	"net/http"
)

func RegisterRoutes() {

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/students", getStudents)
}

// 🏠 Home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Lab 7 API 🚀")
}