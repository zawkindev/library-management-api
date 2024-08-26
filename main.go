package main

import (
	h "library-management-api/handler"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	mainMux := http.NewServeMux()

	booksMux := http.NewServeMux()
	booksMux.HandleFunc("/", h.BookHandler)

	mainMux.HandleFunc("/", h.HomeHandler)
	mainMux.HandleFunc("/books", h.BooksHandler)
	mainMux.Handle("/books/", http.StripPrefix("/books", booksMux))

	http.ListenAndServe(":8080", mainMux)
}
