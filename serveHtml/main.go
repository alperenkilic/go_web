package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/:name", func(c *fiber.Ctx) error {
		if c.Params("name") == "anasayfa" {
			return c.SendFile("./anasayfa.html")
		}
		if c.Params("name") == "blog" {
			return c.SendFile("./blog.html")
		}
		def := fmt.Sprintf("Hello %v", c.Params("name"))
		return c.Send([]byte(def))
	})
	app.Listen(":3000")
}
