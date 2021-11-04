package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/controllers"
)

func Init(app *fiber.App) {

	app.Get("/api/users/currentuser", controllers.GetCurrentUser)
	app.Post("/api/users/signin", controllers.Signin)
	app.Post("/api/users/signup", controllers.Signup)
	app.Post("/api/users/signout", controllers.Signout)

}
