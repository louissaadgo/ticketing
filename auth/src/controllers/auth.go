package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
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

	database.DB.InsertOne(context.TODO(), user)

	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: generateJWT(24, user.Email),
	}
	c.Cookie(&cookie)

	return c.JSON(user)
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Hi there signout")
}

func generateJWT(id int, email string) string {

	signingKey := []byte("12345")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["email"] = email
	claims["iss"] = "auth"
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, _ := token.SignedString(signingKey)

	return tokenString
}
