package controllers

import "github.com/gofiber/fiber/v2"

func GetOrders(c *fiber.Ctx) error {
	return c.SendString("HI")
}

func GetOrderByID(c *fiber.Ctx) error {
	return c.SendString("HI")
}

func CreateOrder(c *fiber.Ctx) error {
	return c.SendString("HI")
}

func DeleteOrder(c *fiber.Ctx) error {
	return c.SendString("HI")
}
