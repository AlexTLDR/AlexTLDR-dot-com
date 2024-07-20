package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./src")

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	app.Listen(":3000")
}
