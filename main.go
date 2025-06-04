package main

import (
	"log"
	"net/http"
	"your_module_name/handlers"
)

func main() {
	http.HandleFunc("/books", handlers.CreateBook) // тільки POST поки
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
