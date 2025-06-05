package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sourire-lanuit/lab5/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBookHandler).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
