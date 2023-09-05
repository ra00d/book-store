package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/app/handlers"
)

func AuthController(app *fiber.App) {
	// ! DO NOT ADD NIL HANDLER
	router := app.Group("auth")
	router.Post("login", handlers.Login)
}
