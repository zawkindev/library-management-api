package main

import (
	"library-management-api/database"
	h "library-management-api/handler"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	database.ConnectDB()
	// Get the underlying sql.DB object
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Failed to retrieve the database object: %v", err)
	}

	// Ensure the database connection is closed when the application exits
	defer sqlDB.Close()

	mainMux := http.NewServeMux()

	mainMux.HandleFunc("/", h.HomeHandler)
	mainMux.HandleFunc("/books/", h.BooksHandler)

	http.ListenAndServe(":8080", mainMux)
}
