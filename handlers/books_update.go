package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sourire-lanuit/lab5/models"
)


var books = []models.Book{
    {ID: "1", Title: "Book One", Author: "Author A", Year: 2001},
    {ID: "2", Title: "Book Two", Author: "Author B", Year: 2005},
}

func UpdateBook(c *gin.Context) {
    id := c.Param("id")

    var updatedBook models.Book
    if err := c.ShouldBindJSON(&updatedBook); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    for i, b := range books {
        if b.ID == id {
            // Оновлюємо поля книги, зберігаємо ідентифікатор незмінним
            updatedBook.ID = id
            books[i] = updatedBook

            c.JSON(http.StatusOK, updatedBook)
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
