package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/tickets/src/database"
	"github.com/louissaadgo/ticketing/tickets/src/models"
)

func CreateTicket(c *fiber.Ctx) error {

	ticket := models.Ticket{}

	//Validating received data type
	if err := c.BodyParser(&ticket); err != nil {
		validationError := models.Error{}
		validationError.Message = "Received invalid data type"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, validationError)
		c.Status(400)
		return c.JSON(errorResponse)
	}

	//Validating received data
	if isValid := ticket.Validate(); !isValid {
		validationError := models.Error{}
		validationError.Message = "Received invalid data"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, validationError)
		c.Status(400)
		return c.JSON(errorResponse)
	}

	//Inserting ticket into DB
	//Handle db error later
	database.DB.InsertOne(context.TODO(), ticket)

	return c.JSON(ticket)
}
