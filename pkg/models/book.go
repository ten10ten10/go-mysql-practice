package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ten10ten10/go-mysql-practice/pkg/config"
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
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) *gorm.DB {
	var book Book
	db := db.Where("ID=?", Id).Delete(book)
	return db
}

func (b *Book) UpdateBook(Id int64) *Book {
	db.Model(&b).Where("ID=?", Id).Updates(b)
	return b
}
