package handler

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome man!")
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Error(w, "book ID is not specified!", http.StatusBadRequest)
		return
	}

	id := r.URL.Path[1:]

	switch r.Method {
	case http.MethodGet:
		fmt.Fprintln(w, "book ID: ", id)
	case http.MethodPut:
		updateBook(w, r, id)
	case http.MethodDelete:
		deleteBook(w, r, id)
	}
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getBooks(w, r)
	case http.MethodPost:
		createBook(w, r)
	}
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "list of books")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	fmt.Println("BOOK CREATED")
}

func updateBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintln(w, "UPDATED book ID: ", id)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintln(w, "DELETED book ID: ", id)
}
