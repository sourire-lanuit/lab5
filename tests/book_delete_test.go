package handlers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
    "github.com/sourire-lanuit/lab5/handlers"
	"github.com/sourire-lanuit/lab5/models"
)

func addBookForDelete(t *testing.T) models.Book {
	book := models.Book{
		Title:  "Delete Me",
		Author: "Bye Author",
		Pages:  99,
	}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest("POST", "/books", bytes.NewReader(body))
	w := httptest.NewRecorder()
	handlers.CreateBook(w, req)

	var created models.Book
	json.NewDecoder(w.Body).Decode(&created)
	return created
}

func TestDeleteBook(t *testing.T) {
	book := addBookForDelete(t)

	req := httptest.NewRequest("DELETE", "/books/"+book.ID, nil)
	w := httptest.NewRecorder()
	handlers.DeleteBook(w, req)

	if w.Result().StatusCode != http.StatusNoContent {
		t.Errorf("Expected 204 No Content, got %v", w.Result().Status)
	}

	reqCheck := httptest.NewRequest("GET", "/books/"+book.ID, nil)
	wCheck := httptest.NewRecorder()
	handlers.GetBookHandler(wCheck, reqCheck)

	if wCheck.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Book was not deleted")
	}
}
