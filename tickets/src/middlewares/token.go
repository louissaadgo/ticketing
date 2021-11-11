package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/tickets/src/token"
)

func IsValidPasetoToken(c *fiber.Ctx) error {

	tokenString := c.Cookies("token")
	payload, isValid := token.VerifyPasetoToken(tokenString)

	if !isValid {
		c.Status(fiber.StatusUnauthorized)
		return c.SendString("You must be authenticated")
	}

	c.Set("CurrentUser", payload.Email)

	return c.Next()
}
