package controllers

import "github.com/gofiber/fiber/v2"

func GetCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Hi there CurrentUser")
}

func Signin(c *fiber.Ctx) error {
	return c.SendString("Hi there signin")
}

func Signup(c *fiber.Ctx) error {
	return c.SendString("Hi there signup")
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Hi there signout")
}
