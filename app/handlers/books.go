package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/app/helpers"
	"github.com/ra00d/book_store/app/models"
	"gorm.io/gorm"
)

func GetBooksHandler(c *fiber.Ctx) error {
	books := []models.Book{}
	db := helpers.GetDb(c)
	db.Find(&books)

	return c.JSON(books)
}

func CreateBookHandler(c *fiber.Ctx) error {
	book := &models.Book{}
	// err := c.BodyParser(body)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"statusCode": 500,
	// 		"message":    "unable to parser the data",
	// 	})
	// }
	image, err := c.FormFile("cover")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"statusCode": 500,
			"message":    "unable to parser the cover image",
		})
	}
	pdf, err := c.FormFile("file")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"statusCode": 500,
			"message":    "unable to parser the cover image",
		})
	}
	book.Image = fmt.Sprintf("uploads/covers/%s", image.Filename)
	if err := c.SaveFile(image, fmt.Sprintf("./storage/uploads/covers/%s", image.Filename)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"statusCode": 500,
			"message":    "unable to save the cover image",
		})
	}
	book.File = fmt.Sprintf("uploads/pdfs/%s", pdf.Filename)
	if err := c.SaveFile(pdf, fmt.Sprintf("./storage/uploads/pdfs/%s", pdf.Filename)); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"statusCode": 500,
			"message":    "unable to save the cover image",
		})
	}

	book.Name = c.FormValue("name")
	book.Auther = c.FormValue("auther")
	book.Edition = c.FormValue("edition")
	db := helpers.GetDb(c)
	res := db.Create(book)
	if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
		return c.Status(422).JSON(fiber.Map{
			"message": "invalid email address used befor",
		})
	}

	return c.SendStatus(fiber.StatusAccepted)
}

func GetBookHandler(c *fiber.Ctx) error {
	book := models.Book{}
	db := helpers.GetDb(c)
	res := db.First(&book, "id=?", c.Params("id", "0"))
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"statusCode": "404",
			"message":    "Book is not found",
		})
	}

	return c.JSON(book)
}

func DeleteBookHandler(c *fiber.Ctx) error {
	book := models.Book{}
	db := helpers.GetDb(c)
	res := db.Delete(&book, "id=?", c.Params("id", "0"))
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"statusCode": "404",
			"message":    "Book is not found",
		})
	}

	if res.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"statusCode": "500",
			"message":    "internal server error",
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}
