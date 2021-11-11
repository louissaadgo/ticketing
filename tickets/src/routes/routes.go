package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/tickets/src/controllers"
)

func Init(app *fiber.App) {

	app.Post("/api/tickets", controllers.CreateTicket)

}
