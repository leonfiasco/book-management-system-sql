package model

import (
	"github.com/leonfiasco/go-bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) // Automatically migrate the schema of the database
}

// CreateBook inserts a new book record into the database.
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

// GetAllBooks retrieves all book records from the database.
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

// GetBookByID retrieves a book record by its ID.
func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(&getBook)
	return &getBook, db
}

// DeleteBook deletes a book record by its ID.
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}
