package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/routes"
)

func main() {
	app := fiber.New()

	routes.Init(app)

	app.Listen(":3000")
}
