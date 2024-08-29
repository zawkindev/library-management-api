package handler

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"library-management-api/database"
	"library-management-api/model"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome man!")
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		getBooks(w, r)
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
	var books []model.Book

	if err := database.DB.Find(&books).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)

}

func getBook(w http.ResponseWriter, r *http.Request, id string) {
	var book model.Book

	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = database.DB.Create(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintln(w, "UPDATED book ID: ", id)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprintln(w, "DELETED book ID: ", id)
}

func strToHash(input string) string {
	hash := md5.New()
	io.WriteString(hash, input)
	hashBytes := hash.Sum(nil)

	return fmt.Sprintf("%x", hashBytes)
}
