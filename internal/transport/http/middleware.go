package http

import "github.com/gofiber/fiber/v2"

func CheckAgeMiddleware(c *fiber.Ctx) error {
	age := c.Locals("age").(int)
	if age < 18 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Пользователь несовершеннолетний"})
	}
	return c.Next()
}
