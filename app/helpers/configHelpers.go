package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ra00d/book_store/config"
	"gorm.io/gorm"
)

func GetDb(c *fiber.Ctx) *gorm.DB {
	config := c.Locals("app").(*config.AppConfig)

	return config.Db
}
