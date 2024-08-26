package handler

import (
	"encoding/json"
	"fmt"
	"library-management-api/database"
	"library-management-api/model"
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
		getBook(w, r, id)
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
	rows, err := database.DB.Query("SELECT id, title, author, year, genre, isbn FROM books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var books []model.Book
	for rows.Next() {
		var book model.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.Genre, &book.ISBN)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintln(w, "book ID: ", id)
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
