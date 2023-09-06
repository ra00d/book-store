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
	res := db.Where("email=?", body.Email).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.Status(404).JSON(fiber.Map{
			"Select":  true,
			"message": "wrong password or email",
		})
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"pass":      body.Password,
			"user_pass": user.Password,
			"error":     err.Error(),
			"message":   "wrong password or email",
		})
	}
	return c.Status(fiber.StatusAccepted).JSON(user)
}

func SignUp(c *fiber.Ctx) error {
	db := helpers.GetDb(c)
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   err,
			"message": "invalid input",
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(403).JSON(fiber.Map{
			"user_pass": user.Password,
			"error":     err.Error(),
			"message":   "wrong password or email",
		})
	}
	user.Password = string(hash)
	user.Role.Name = "user"
	// err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(string(hash)))
	// if err != nil {
	// 	return c.Status(403).JSON(fiber.Map{
	// 		"user_pass": user.Password,
	// 		"error":     err.Error(),
	// 		"message":   "wrong password or email",
	// 	})
	// }
	//
	res := db.Create(&user)
	if errors.Is(res.Error, gorm.ErrDuplicatedKey) {
		return c.Status(404).JSON(fiber.Map{
			"message": "invalid email address used befor",
		})
	}
	return c.SendStatus(fiber.StatusCreated)
}
