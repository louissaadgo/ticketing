package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/bus"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/models"
	"github.com/louissaadgo/ticketing/auth/src/routes"
	"github.com/nats-io/stan.go"
)

func main() {
	//Add testing later
	bus.CreateSTANConnection()

	go bus.CreateSTANListener(bus.TicketCreatedEvent, "auth", func(m *stan.Msg) {
		ticket := models.Ticket{}
		json.Unmarshal(m.Data, &ticket)
		fmt.Println(ticket)
		m.Ack()
	})

	app := fiber.New()

	routes.Init(app)

	database.Connect()

	app.Listen(":3000")
}
