package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Constants
	const (
		PORT = ":3000"
	)

	// Set up app
	app := fiber.New()

	// Define a route
	app.Get("/", func(c *fiber.Ctx) error {
		// Render the "index.html" template
		return c.Render("../templates/home.html", fiber.Map{
			"Name": "Jowi",
		})
	})

	// Listen on the port
	app.Listen(PORT)
}
