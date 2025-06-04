package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"your_module_name/handlers"
	"your_module_name/models"
)

func addBookForRead(t *testing.T) models.Book {
	book := models.Book{
		Title:  "ReadBook",
		Author: "Reader",
		Pages:  101,
	}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handlers.CreateBook(w, req)

	var created models.Book
	json.NewDecoder(w.Body).Decode(&created)
	return created
}

func TestGetBook(t *testing.T) {
	book := addBookForRead(t)

	req := httptest.NewRequest("GET", "/books/"+book.ID, nil)
	w := httptest.NewRecorder()

	handlers.GetBook(w, req)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Result().StatusCode)
	}
}

func TestListBooks(t *testing.T) {
	addBookForRead(t)

	req := httptest.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()

	handlers.ListBooks(w, req)
	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v", w.Result().Status)
	}

	var books []models.Book
	if err := json.NewDecoder(w.Body).Decode(&books); err != nil {
		t.Errorf("Decode failed: %v", err)
	}

	if len(books) == 0 {
		t.Errorf("Expected at least one book")
	}
}
