package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/models"
	"github.com/louissaadgo/ticketing/auth/src/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCurrentUser(c *fiber.Ctx) error {
	email := c.GetRespHeader("CurrentUser")
	return c.JSON(fiber.Map{
		"currentUser": email,
	})
}

func Signin(c *fiber.Ctx) error {

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

	//Cheking if user already exists
	userCheck := models.User{}
	err := database.DB.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&userCheck)
	if err == mongo.ErrNoDocuments {
		queryError := models.Error{
			Message: "User not found",
		}
		queryErrorResponse := models.ErrorResponse{}
		queryErrorResponse.Errors = append(queryErrorResponse.Errors, queryError)
		c.Status(400)
		return c.JSON(queryErrorResponse)
	}

	//Checking credentials
	isValid := user.CompareHashAndPassword(userCheck.Password)
	if !isValid {
		validationError := models.Error{}
		validationError.Message = "Invalid password"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, validationError)
		c.Status(400)
		return c.JSON(errorResponse)
	}

	//Creating and sending a paseto cookie
	token, err := token.GeneratePasetoToken(user.Email)
	if err != nil {
		tokenError := models.Error{}
		tokenError.Message = "Token creation failed"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, tokenError)
		c.Status(400)
		return c.JSON(errorResponse)
	}
	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	c.Cookie(&cookie)

	user.Password = ""
	return c.JSON(user)
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

	//Cheking if user already exists
	query := database.DB.FindOne(context.TODO(), bson.M{"email": user.Email})
	if query.Err() != mongo.ErrNoDocuments {
		queryError := models.Error{
			Message: "Email already in use",
		}
		queryErrorResponse := models.ErrorResponse{}
		queryErrorResponse.Errors = append(queryErrorResponse.Errors, queryError)
		c.Status(400)
		return c.JSON(queryErrorResponse)
	}

	//Hashing the password
	if err := user.HashPassword(); !err {
		hashingError := models.Error{}
		hashingError.Message = "Unable to hash password"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, hashingError)
		c.Status(400)
		return c.JSON(errorResponse)
	}

	//Inserting user into the db
	database.DB.InsertOne(context.TODO(), user)

	//Creating and sending a paseto cookie
	token, err := token.GeneratePasetoToken(user.Email)
	if err != nil {
		tokenError := models.Error{}
		tokenError.Message = "Token creation failed"
		errorResponse := models.ErrorResponse{}
		errorResponse.Errors = append(errorResponse.Errors, tokenError)
		c.Status(400)
		return c.JSON(errorResponse)
	}
	cookie := fiber.Cookie{
		Name:  "token",
		Value: token,
	}
	c.Cookie(&cookie)

	user.Password = ""

	return c.JSON(user)
}

func Signout(c *fiber.Ctx) error {

	cookie := fiber.Cookie{
		Name:    "token",
		Expires: time.Now().Add(-time.Minute),
	}
	c.Cookie(&cookie)

	return c.SendString("Bye Bye")
}
