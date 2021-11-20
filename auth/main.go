package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	go app.Listen(":3000")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	if err := app.Shutdown(); err != nil {
		fmt.Println("Graceful Shutdown failed: ", err)
	}
	fmt.Println("Graceful Shutdown success")

	if err := bus.SC.Close(); err != nil {
		fmt.Println("Graceful Stan Shutdown failed: ", err)
	}
	fmt.Println("Graceful Stan Shutdown success")
}
