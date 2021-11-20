package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/bus"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/routes"
)

func main() {
	//Add testing later
	bus.CreateSTANConnection()

	app := fiber.New()

	routes.Init(app)

	database.Connect()

	app.Listen(":3000")
}
