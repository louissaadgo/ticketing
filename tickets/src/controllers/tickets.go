package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/louissaadgo/ticketing/tickets/src/database"
	"github.com/louissaadgo/ticketing/tickets/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllTickets(c *fiber.Ctx) error {

	//Handle error later
	cur, _ := database.DB.Find(context.TODO(), bson.D{})
	response := models.ManyTicketsResponse{}

	for cur.Next(context.TODO()) {
		ticket := models.Ticket{}
		cur.Decode(&ticket)
		response.Tickets = append(response.Tickets, ticket)
	}

	return c.JSON(response)
}

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

	//Adding userid and ticketid to ticket
	ticket.UserID = c.GetRespHeader("CurrentUser")
	ticket.TicketID = uuid.New().String()

	//Inserting ticket into DB
	//Handle db error later
	database.DB.InsertOne(context.TODO(), ticket)

	return c.JSON(ticket)
}

func RetreiveTicket(c *fiber.Ctx) error {

	ticketID := c.Params("id")
	ticket := models.Ticket{}

	err := database.DB.FindOne(context.TODO(), bson.M{"ticketid": ticketID}).Decode(&ticket)
	if err == mongo.ErrNoDocuments {
		queryError := models.Error{
			Message: "Ticket not found",
		}
		queryErrorResponse := models.ErrorResponse{}
		queryErrorResponse.Errors = append(queryErrorResponse.Errors, queryError)
		c.Status(404)
		return c.JSON(queryErrorResponse)
	}

	return c.JSON(ticket)
}
