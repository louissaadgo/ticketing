package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/api/users/currentuser", func(c *fiber.Ctx) error {
		return c.SendString("Hi there!")
	})

	app.Listen(":3000")
}
