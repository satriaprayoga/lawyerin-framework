package data

import "gorm.io/gorm"

var db *gorm.DB

type Store struct {
	Book Book
}

func New(conn *gorm.DB) *Store {
	db = conn
	autoMigrate()
	return &Store{Book: Book{}}
}

func autoMigrate() {
	db.AutoMigrate(Book{})
}
