package middlewares

import "github.com/gofiber/fiber/v2"

func FileSizeLimit(maxSize int64) fiber.Handler{
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"error":"No File Provided",
			})
		} 
		if file.Size > maxSize {
			return c.Status(413).JSON(fiber.Map{
				"error":"File size exceeds limit of 5mb",
			})
		}
		return c.Next()
	}
}