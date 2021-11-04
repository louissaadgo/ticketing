package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/models"
)

func GetCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Hi there CurrentUser")
}

func Signin(c *fiber.Ctx) error {
	return c.SendString("Hi there signin")
}

func Signup(c *fiber.Ctx) error {

	//Validating received data type
	user := models.User{}
	if err := c.BodyParser(&user); err != nil {
		c.Status(400)
		return c.SendString("Received invalid data type")
	}

	//Validating received data
	if isValid := user.ValidateUserModel(); !isValid {
		c.Status(400)
		return c.SendString("Received invalid data")
	}

	return c.SendString("Successful signup")
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Hi there signout")
}
