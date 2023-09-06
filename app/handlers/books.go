package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/app/helpers"
	"github.com/ra00d/book_store/app/models"
)

func GetBooksHandler(c *fiber.Ctx) error {
	books := []models.Book{}
	db := helpers.GetDb(c)
	db.Find(&books)

	return c.JSON(books)
}
