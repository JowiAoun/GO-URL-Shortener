package main

import (
  "github.com/gofiber/fiber/v2"
)

func main() {
  // Create a new Fiber app
  app := fiber.New()

  // Define a route
  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })

  // Start the server
  app.Listen(":3000")
}
