package server

import (
  "encoding/json"
  "net/http"
  "strings"

  "github.com/sourire-lanuit/lab5/handlers"
)

type router struct{}

func NewRouter() http.Handler {
  return &router{}
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  path := req.URL.Path
  method := req.Method

  switch {
	case path == "/health" && method == http.MethodGet:
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"status": "ok"})

  case path == "/books" && method == http.MethodPost:
    handlers.CreateBook(w, req)

  case strings.HasPrefix(path, "/books/"):
    id := strings.TrimPrefix(path, "/books/")
    if id == "" {
      http.NotFound(w, req)
      return
    }

    switch method {
    case http.MethodGet:
      handlers.GetBookHandler(w, req)
    case http.MethodPut:
      handlers.UpdateBook(w, req)
    case http.MethodDelete:
      handlers.DeleteBook(w, req)
    default:
      http.NotFound(w, req)
    }

  default:
    http.NotFound(w, req)
  }
}