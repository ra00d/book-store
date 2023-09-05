package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/app/helpers"
	"github.com/ra00d/book_store/app/models"
	"github.com/ra00d/book_store/app/requests"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	db := helpers.GetDb(c)
	body := requests.LoginRequestBody{}
	user := new(models.User)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err,
			"message": "invalid input",
		})
	}
	res := db.Select("password").Where("email=?", body.Email).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"message": "wrong password or email",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"message": "wrong password or email",
		})
	}
	return c.SendStatus(fiber.StatusAccepted)
}
