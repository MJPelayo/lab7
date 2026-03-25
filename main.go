package main

// Import required packages
import (
	"database/sql" // allows Go to talk to databases
	"fmt"          // for printing messages
	"log"          // for logging errors
	"net/http"     // for creating a web server

	_ "github.com/lib/pq" // PostgreSQL driver
)

// 🌍 Global database variable (accessible in all files)
var DB *sql.DB

func main() {

	// 🔗 Database connection string (DSN)
	dsn := "postgres://go_user:go123@localhost:5432/myapp_db?sslmode=disable"

	var err error

	// 🛠️ Open database connection
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error opening database:", err)
	}

	// 🧪 Check if database is reachable
	err = DB.Ping()
	if err != nil {
		log.Fatal("❌ Database not reachable:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")

	// 🌐 ROUTES (API endpoints)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/students", getStudents) // 👈 NEW ROUTE

	// 🚀 Start server
	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// 🏠 Home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Lab 7 API 🚀")
}