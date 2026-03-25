package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// 🌍 Global DB variable
var DB *sql.DB

// 🔗 Initialize DB connection
func InitDB() {

	dsn := "postgres://go_user:go123@localhost:5432/myapp_db?sslmode=disable"

	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error opening database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database not reachable:", err)
	}

	log.Println("✅ Connected to PostgreSQL!")
}