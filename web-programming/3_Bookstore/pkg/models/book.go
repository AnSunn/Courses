package models

import (
	"github.com/AnSunn/3_Bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model        //Model is a struct which declare ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `gorm:""json:"name"`
	Author     string `json:"author"`
	Publicaion string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //creates table books
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b) // returns true if the current value's primary key is unset, meaning it's a new record
	db.Create(&b)   //insert the value into db
	return b
}
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook) //SELECT * FROM books WHERE ID = Id
	//db := db.First(&getBook, Id)
	return &getBook, db
}
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID=?", Id).Delete(book) //delete from books where ID = Id
	return book
}
