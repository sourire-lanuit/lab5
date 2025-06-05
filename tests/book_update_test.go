package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/sourire-lanuit/lab5/handlers"
    "github.com/sourire-lanuit/lab5/models"
)

func setupRouter() *gin.Engine {
    r := gin.Default()
    r.PUT("/books/:id", handlers.UpdateBook)
    return r
}

func TestUpdateBook_Success(t *testing.T) {
    router := setupRouter()

    updated := models.Book{
        Title:  "Updated Title",
        Author: "Updated Author",
        Year:   2023,
    }
    jsonValue, _ := json.Marshal(updated)

    req, _ := http.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    var respBook models.Book
    err := json.Unmarshal(w.Body.Bytes(), &respBook)
    assert.NoError(t, err)

    assert.Equal(t, "1", respBook.ID)
    assert.Equal(t, updated.Title, respBook.Title)
    assert.Equal(t, updated.Author, respBook.Author)
    assert.Equal(t, updated.Year, respBook.Year)
}

func TestUpdateBook_NotFound(t *testing.T) {
    router := setupRouter()

    updated := models.Book{
        Title:  "Book",
        Author: "No One",
        Year:   2025,
    }
    jsonValue, _ := json.Marshal(updated)

    req, _ := http.NewRequest(http.MethodPut, "/books/999", bytes.NewBuffer(jsonValue))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusNotFound, w.Code)
}
