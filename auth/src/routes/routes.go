package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/controllers"
	"github.com/louissaadgo/ticketing/auth/src/middlewares"
)

func Init(app *fiber.App) {

	app.Post("/api/users/signin", controllers.Signin)
	app.Post("/api/users/signup", controllers.Signup)
	app.Post("/api/users/signout", controllers.Signout)
	app.Use(middlewares.IsValidPasetoToken)
	app.Get("/api/users/currentuser", controllers.GetCurrentUser)

}
