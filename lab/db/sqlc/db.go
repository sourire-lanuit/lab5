package db

import (
  "log"

  "github.com/sourire-lanuit/lab5/models"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
  var err error
  DB, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
  if err != nil {
    log.Fatalf("failed to connect database: %v", err)
  }

  err = DB.AutoMigrate(&models.Book{})
  if err != nil {
    log.Fatalf("failed to migrate database: %v", err)
  }
}