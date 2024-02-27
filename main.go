package main

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		port = "3000"
	}

	app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, world!")
    })

	app.Listen(":" + port)
}
