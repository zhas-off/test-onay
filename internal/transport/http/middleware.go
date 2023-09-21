package http

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func CheckAgeMiddleware(c *fiber.Ctx) error {
	if c.Method() == fiber.MethodPost || c.Method() == fiber.MethodPut {
		var postUserReq PostUserRequest

		if err := c.BodyParser(&postUserReq); err != nil {
			log.Error(err)
			return c.Status(fiber.StatusBadRequest).SendString("Bad Request")
		}

		if postUserReq.Age < 18 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Пользователь несовершеннолетний"})
		}

		if postUserReq.Age > 100 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "У пользователя огромный возраст"})
		}
	}
	return c.Next()
}
