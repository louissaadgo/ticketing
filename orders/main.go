package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber"
	"github.com/louissaadgo/ticketing/orders/src/bus"
	"github.com/louissaadgo/ticketing/orders/src/database"
)

func main() {
	//Add testing later
	bus.CreateSTANConnection()

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

	// if err := bus.SC.Close(); err != nil {
	// 	fmt.Println("Graceful Stan Shutdown failed: ", err)
	// }
	// fmt.Println("Graceful Stan Shutdown success")
}
