package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./templates")

	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	// Log a message when the server starts
	err := app.Listen(":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		fmt.Println("Server started on port 3000")
	}
}
