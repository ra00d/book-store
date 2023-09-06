package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/app/handlers"
)

func BooksController(app *fiber.App) {
	// ! DO NOT ADD NIL HANDLER
	router := app.Group("books")
	router.Get("", handlers.GetBooksHandler)
}
