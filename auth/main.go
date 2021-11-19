package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/bus"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/routes"
	"github.com/nats-io/stan.go"
)

func main() {
	//Add testing later
	bus.CreateSTANConnection()

	go bus.CreateSTANListener("auth:created-user", "auth-service", func(m *stan.Msg) {
		fmt.Printf("New message: %v", m.Data)
		m.Ack()
	})

	app := fiber.New()

	routes.Init(app)

	database.Connect()

	app.Listen(":3000")
}
