package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/django/v3"
	"github.com/joho/godotenv"

	"github.com/ra00d/book_store/config"
)

type Test struct {
	k1 string
	k2 string
}

func main() {
	godotenv.Load()
	engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
		// EnablePrintRoutes: true,
		PassLocalsToViews: true,
	})
	app.Static("", "./public")
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.JSON("test")
	})
	appConfig := config.AppConfig{}

	appConfig.Init()
	app.Use(func(c *fiber.Ctx) error {
		// appConfig.Db.AutoMigrate(&books.Book{})
		c.Locals("app", &appConfig)
		// c.Locals("user", &auth.User{})
		return c.Next()
	})

	app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
