package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/orders/src/controllers"
	"github.com/louissaadgo/ticketing/orders/src/middlewares"
)

func Init(app *fiber.App) {
	app.Use(middlewares.IsValidPasetoToken)
	app.Get("/api/orders", controllers.GetOrders)
	app.Get("/api/orders/:id", controllers.GetOrderByID)
	app.Post("/api/orders", controllers.CreateOrder)
	app.Delete("/api/orders/:id", controllers.DeleteOrder)
}
