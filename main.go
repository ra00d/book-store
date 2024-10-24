package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"

	"github.com/ra00d/book_store/app/controllers"
	// "github.com/ra00d/book_store/app/models"
	"github.com/ra00d/book_store/config"
	// "github.com/ra00d/book_store/database/seeders"
)

type Test struct {
	k1 string
	k2 string
}

func main() {
	godotenv.Load()
	// engine := django.New("./views", ".html")
	app := fiber.New(fiber.Config{
		// Views:             engine,
		// ViewsLayout:       "layouts/main",
		// EnablePrintRoutes: true,
		// PassLocalsToViews: true,
	})
	app.Static("", "./public")
	app.Use(filesystem.New(filesystem.Config{Root: http.Dir("./storage/")}))
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
	// TODO make this in separate files not in the main function
	// appConfig.Db.AutoMigrate(&models.Permission{}, &models.Role{}, &models.User{}, &models.Like{}, &models.Comment{}, &models.Book{})
	// seeders.UserSeeder(appConfig.Db)
	// seeders.BooksSeeder(appConfig.Db)
	controllers.AuthController(app)
	controllers.BooksController(app)
	app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
