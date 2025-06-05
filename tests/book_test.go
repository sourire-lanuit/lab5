package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"github.com/sourire-lanuit/lab5/handlers"
)

func TestCreateBook_Success(t *testing.T) {
	body := `{"title": "Test Book", "author": "Bertolt Brecht", "pages": 124}`
	req := httptest.NewRequest("POST", "/books", strings.NewReader(body))
	rec := httptest.NewRecorder()

	handlers.CreateBook(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %d", rec.Code)
	}
	if !strings.Contains(rec.Body.String(), "Test Book") {
		t.Errorf("Expected body to contain book title")
	}
}

func TestCreateBook_Invalid(t *testing.T) {
	body := `{"title": "", "author": "", "pages": 0}`
	req := httptest.NewRequest("POST", "/books", strings.NewReader(body))
	rec := httptest.NewRecorder()

	handlers.CreateBook(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", rec.Code)
	}
}
