package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")

	// Test handler
	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("App running")
	})

	log.Fatal(app.Listen(":5000"))
}
