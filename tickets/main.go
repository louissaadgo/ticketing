package main

import "github.com/gofiber/fiber/v2"

func main() {
	//Add testing later

	app := fiber.New()

	app.Listen(":3000")
}
