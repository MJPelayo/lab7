package main

import (
	"log"
	"net/http"
)

func main() {

	db := openDB()

	app := &application{
		db: db,
	}

	mux := http.NewServeMux()
	app.routes(mux)

	log.Println("🚀 Server running on :4000")
	log.Fatal(http.ListenAndServe(":4000", mux))
}
