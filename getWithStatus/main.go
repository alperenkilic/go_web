package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/lookstatus", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusBadRequest).Send([]byte("hello but look at the status "))
		// google developer toolkit -> network -> all -> Summary
	})
	app.Listen(":3000")
}
