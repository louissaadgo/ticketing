package controllers

import "github.com/gofiber/fiber/v2"

func CreateTicket(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
