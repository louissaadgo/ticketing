package controllers

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/louissaadgo/ticketing/auth/src/database"
	"github.com/louissaadgo/ticketing/auth/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCurrentUser(c *fiber.Ctx) error {

	token := c.Cookies("jwt")
	valid, email := validateJWT(token)
	if !valid {
		return c.JSON(fiber.Map{
			"currentUser": nil,
		})
	}
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

	//Creating and sending a jwt cookie
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: generateJWT(user.Email),
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

	//Creating and sending a jwt cookie
	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: generateJWT(user.Email),
	}
	c.Cookie(&cookie)

	user.Password = ""

	return c.JSON(user)
}

func Signout(c *fiber.Ctx) error {
	return c.SendString("Hi there signout")
}

func generateJWT(email string) string {

	//move env variable checking to when the app starts
	signingKey := []byte(os.Getenv("JWT_KEY"))

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["iss"] = "auth"
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	tokenString, _ := token.SignedString(signingKey)

	return tokenString
}

func validateJWT(token string) (bool, string) {
	//Migrate to paseto
	newToken, _ := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if _, ok := newToken.Claims.(jwt.MapClaims); ok && newToken.Valid {
		return true, fmt.Sprintf("%v", newToken.Claims.(jwt.MapClaims)["email"])
	} else {
		return false, ""
	}
}
