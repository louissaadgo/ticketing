package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/routes"
)

func main() {

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

	// if err := bus.SC.Close(); err != nil {
	// 	fmt.Println("Graceful Stan Shutdown failed: ", err)
	// }
	// fmt.Println("Graceful Stan Shutdown success")
}
