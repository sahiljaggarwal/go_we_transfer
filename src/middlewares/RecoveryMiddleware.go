package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func RecoveryMiddleware(c *fiber.Ctx) error {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
			c.Status(500).JSON(fiber.Map{
				"status":  "ERROR",
				"message": "Internal Server Error",
			})
		}
	}()

	return c.Next()
}