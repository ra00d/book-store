package middlewares

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/gofiber/fiber/v2"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(body interface{}) fiber.Handler {
	errors := make(map[string]string)
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validate = validator.New()
	en_translation.RegisterDefaultTranslations(validate, trans)
	return func(c *fiber.Ctx) error {
		err := c.BodyParser(&body)
		if err != nil {
			return c.Status(422).JSON(fiber.Map{
				"message": "unproccesable entity",
				"code":    "422",
			})
		}
		fmt.Println(body)
		err = validate.Struct(body)
		if err != nil {
			errs := err.(validator.ValidationErrors)
			if errs != nil {
				for _, e := range errs {
					errors[e.Field()] = e.Translate(trans)
				}
			}
			return c.Status(422).JSON(fiber.Map{
				"message": "some fields are invalid",
				"errors":  errors,
			})
		}
		return c.Next()
	}
}
