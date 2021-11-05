package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCurrentUser(c *fiber.Ctx) error {
	return c.SendString("Hi there CurrentUser")
}

func Signin(c *fiber.Ctx) error {
	return c.SendString("Hi there signin")
}

func Signup(c *fiber.Ctx) error {

	user := models.User{}

	//Validating received data type
	if err := c.BodyParser(&user); err != nil {
		validationError := models.Error{}
		validationError.Message = "Received invalid data type"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, validationError)
		c.Status(400)
		return c.JSON(errorResponse)
	}

	//Validating received data
	if isValid := user.ValidateUserModel(); !isValid {
		validationError := models.Error{}
		validationError.Message = "Received invalid data"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, validationError)
		c.Status(400)
		return c.JSON(errorResponse)
	}

	err := database.DB.FindOne(context.TODO(), bson.D{{"email", user.Email}})
	if err.Err() != mongo.ErrNoDocuments {
		return c.SendString("email in use")
	}

	database.DB.InsertOne(context.TODO(), user)

	return c.SendString("Successful signup")
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Hi there signout")
}
