package main

// Import required packages
import (
	"database/sql" // allows Go to talk to databases
	"fmt"          // for printing messages
	"log"          // for logging errors
	"net/http"     // for creating a web server

	_ "github.com/lib/pq" // PostgreSQL driver (underscore = auto-load)
)

// Global database variable (shared across files)
var DB *sql.DB

func main() {

	// 🔗 Connection string (DSN)
	dsn := "postgres://go_user:go123@localhost:5432/myapp_db?sslmode=disable"

	var err error

	// 🛠️ Open connection to database
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	// 🧪 Test if DB is reachable
	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not reachable:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")

	// 🌐 Start web server
	http.HandleFunc("/", homeHandler)

	fmt.Println("🚀 Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 🏠 Basic route handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Lab 7 API 🚀")
}
