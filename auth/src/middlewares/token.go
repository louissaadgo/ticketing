package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/token"
)

func IsValidPasetoToken(c *fiber.Ctx) error {

	tokenString := c.Cookies("token")
	payload, isValid := token.VerifyPasetoToken(tokenString)

	if !isValid {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"currentUser": nil,
		})
	}

	c.Set("CurrentUser", payload.Email)

	return c.Next()
}
