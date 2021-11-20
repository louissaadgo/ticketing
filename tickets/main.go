package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/tickets/src/bus"
	"github.com/louissaadgo/ticketing/tickets/src/database"
	"github.com/louissaadgo/ticketing/tickets/src/routes"
)

func main() {
	//Add testing later
	bus.CreateSTANConnection()

	app := fiber.New()

	routes.Init(app)

	database.Connect()

	app.Listen(":3000")
}
