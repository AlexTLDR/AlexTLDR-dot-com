package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const port = ":1001"

func main() {
	app := fiber.New()

	// Serve files from the templates directory at the root
	app.Static("/", "./templates")

	// Serve cv.html for the /cv route
	app.Get("/cv", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/cv.html")
	})

	// Serve portfolio.html for the /portfolio route
	app.Get("/portfolio", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/portfolio.html")
	})

	// Serve stuttgart.html for the /stuttgart route
	app.Get("/stuttgart-gophers", func(c *fiber.Ctx) error {
		return c.SendFile("./templates/stuttgart.html")
	})

	// Middleware to redirect undefined routes to /
	app.Use(func(c *fiber.Ctx) error {
		return c.Redirect("/")
	})

	// Log a message when the server starts
	err := app.Listen(port)
	if err != nil {
		fmt.Println("Error starting server:", err)
	} else {
		// Correct the port in the log message to match the Listen port
		fmt.Println("Server started on port", port)
	}
}
