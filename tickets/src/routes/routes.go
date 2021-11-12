package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/tickets/src/controllers"
	"github.com/louissaadgo/ticketing/tickets/src/middlewares"
)

func Init(app *fiber.App) {

	app.Get("/api/tickets/:id", controllers.RetreiveTicket)
	app.Use(middlewares.IsValidPasetoToken)
	app.Post("/api/tickets", controllers.CreateTicket)
}
