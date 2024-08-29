package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"library-management-api/database"
	"library-management-api/model"
	"net/http"

	"gorm.io/gorm"
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
	var book model.Book

	// Find the book by ID first
	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Decode the JSON request body into the book struct (this will overwrite the book fields)
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the book record with the new data
	if err := database.DB.Save(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated book
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func deleteBook(w http.ResponseWriter, r *http.Request, id string) {
	var book model.Book

	// Find the book by ID first to ensure it exists
	if err := database.DB.First(&book, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Delete the book from the database
	if err := database.DB.Delete(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a 204 No Content status to indicate successful deletion
	w.WriteHeader(http.StatusNoContent)
}
