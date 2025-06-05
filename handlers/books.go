package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
	"github.com/gorilla/mux"

	"github.com/sourire-lanuit/lab5/models"
)

var books = make(map[string]models.Book)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &book)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}


	if strings.TrimSpace(book.Title) == "" || len(book.Title) > 255 {
		http.Error(w, "Title is required and must be < 255", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(book.Author) == "" || len(book.Author) > 255 {
		http.Error(w, "Author is required and must be < 255", http.StatusBadRequest)
		return
	}
	if book.Pages <= 0 {
		http.Error(w, "Pages must be > 0", http.StatusBadRequest)
		return
	}


	book.ID = uuid.New().String()
	books[book.ID] = book

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    book, found := books[id]
    if !found {
        http.Error(w, "Book not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := books[id]; !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(books, id)

	w.WriteHeader(http.StatusNoContent)
}

func ListBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bookList := make([]models.Book, 0, len(books))
	for _, b := range books {
		bookList = append(bookList, b)
	}

	json.NewEncoder(w).Encode(bookList)
}
