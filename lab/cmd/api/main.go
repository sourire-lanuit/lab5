package main

import (
  "log"
  "net/http"

  "github.com/sourire-lanuit/lab5/lab/internal/server"

)

func main() {
  r := server.NewRouter()

  log.Println("Server running on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}