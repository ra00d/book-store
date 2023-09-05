package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/app/handlers"
	"github.com/ra00d/book_store/app/middlewares"
	"github.com/ra00d/book_store/app/requests"
)

func AuthController(app *fiber.App) {
	// ! DO NOT ADD NIL HANDLER
	router := app.Group("auth")
	router.Post("login", middlewares.Validate(&requests.LoginRequestBody{}), handlers.Login)
}
