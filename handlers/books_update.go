package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/sourire-lanuit/lab5/models"
)


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	existing, ok := books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	
	var updated models.Book
	body, _ := io.ReadAll(r.Body)
	if err := json.Unmarshal(body, &updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	
	if strings.TrimSpace(updated.Title) == "" || len(updated.Title) > 255 {
		http.Error(w, "Title is required and must be < 255", http.StatusBadRequest)
		return
	}
	if strings.TrimSpace(updated.Author) == "" || len(updated.Author) > 255 {
		http.Error(w, "Author is required and must be < 255", http.StatusBadRequest)
		return
	}
	if updated.Pages <= 0 {
		http.Error(w, "Pages must be > 0", http.StatusBadRequest)
		return
	}

	updated.ID = existing.ID
	books[id] = updated

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}
