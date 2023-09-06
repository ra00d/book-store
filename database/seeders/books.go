package seeders

import (
	"github.com/ra00d/book_store/app/models"
	"gorm.io/gorm"
)

func BooksSeeder(db *gorm.DB) {
	book := models.Book{
		Name:    "book1",
		Auther:  "raad",
		Edition: "third",
		File:    "book1.pdf",
		Image:   "book1.gpj",
	}

	db.Create(&book)
}
