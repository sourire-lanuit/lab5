package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/sourire-lanuit/lab5/handlers"
    "github.com/sourire-lanuit/lab5/models"
)

func TestUpdateBook_Success(t *testing.T) {
    updated := models.Book{
        Title:  "Updated Title",
        Author: "Updated Author",
        Pages:  123,
    }
    jsonValue, _ := json.Marshal(updated)

    req, _ := http.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    r := mux.NewRouter()
    r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusOK {
        t.Fatalf("expected status 200, got %d", w.Code)
    }

    var respBook models.Book
    err := json.Unmarshal(w.Body.Bytes(), &respBook)
    if err != nil {
        t.Fatalf("error decoding response: %v", err)
    }

    if respBook.ID != "1" {
        t.Errorf("expected ID 1, got %s", respBook.ID)
    }
    if respBook.Title != updated.Title {
        t.Errorf("expected Title %q, got %q", updated.Title, respBook.Title)
    }
    if respBook.Author != updated.Author {
        t.Errorf("expected Author %q, got %q", updated.Author, respBook.Author)
    }
    if respBook.Pages != updated.Pages {
        t.Errorf("expected Pages %d, got %d", updated.Pages, respBook.Pages)
    }
}

func TestUpdateBook_NotFound(t *testing.T) {
    updated := models.Book{
        Title:  "Book",
        Author: "No One",
        Pages:  321,
    }
    jsonValue, _ := json.Marshal(updated)

    req, _ := http.NewRequest(http.MethodPut, "/books/999", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    r := mux.NewRouter()
    r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    if w.Code != http.StatusNotFound {
        t.Fatalf("expected status 404, got %d", w.Code)
    }
}
