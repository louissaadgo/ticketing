package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/tickets/src/database"
)

func main() {
	//Add testing later

	app := fiber.New()

	database.Connect()

	app.Listen(":3000")
}
