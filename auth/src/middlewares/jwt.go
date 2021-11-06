package middlewares

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func IsValidJWT(c *fiber.Ctx) error {
	//Move to paseto
	token := c.Cookies("jwt")
	if token == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"currentUser": "",
		})
	}
	newToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if _, ok := newToken.Claims.(jwt.MapClaims); ok && newToken.Valid {
		c.Set("CurrentUser", fmt.Sprintf("%v", newToken.Claims.(jwt.MapClaims)["email"]))
		return c.Next()
	}
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"currentUser": "",
	})
}
