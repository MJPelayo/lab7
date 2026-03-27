package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Git Commit: refactor DB connection to use application struct

func openDB() *sql.DB {

	dsn := "postgres://go_user:go123@localhost:5432/myapp_db?sslmode=disable"

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ DB open error:", err)
	}

	// connection pool settings
	db.SetMaxOpenConns(25)                   // max open connections
	db.SetMaxIdleConns(10)                   // idle connections
	db.SetConnMaxIdleTime(15 * time.Minute)  // idle timeout

	// test connection with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("❌ DB not reachable:", err)
	}

	log.Println("✅ Connected to PostgreSQL")

	return db
}