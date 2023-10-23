package controllers

import (
	"fmt"
	"grincajg/database"
	"grincajg/env"
	"grincajg/models"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// SignUpUser signs up a new user.
// @Summary Sign up a new user
// @Description This endpoint allows a user to sign up by providing necessary information.
// @Tags Users
// @Accept json
// @Produce json
// @Param input body models.SignUpInput true "User signup details" 
// @Success 201 {object} models.UserResponse "User successfully created"
// @Failure 400 {object} models.ValidateErrorResponse "bas request"
// @Failure 409 {object} models.ErrorResponse "Conflict: User with that email already exists"
// @Failure 502 {object} models.ErrorResponse "Bad Gateway: Something bad happened"
// @Router /auth/register [post]
func SignUpUser(c *fiber.Ctx) error {
	var input *models.SignUpInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(input)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	if input.Password != input.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newUser := models.User{
		Name:     input.Name,
		Email:    strings.ToLower(input.Email),
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
}

// SignInUser logs in a user.
// @Summary Log in a user
// @Description This endpoint allows a user to log in by providing necessary information.
// @Tags Users
// @Accept json
// @Produce json
// @Param input body models.SignInInput true "User login details"
// @Success 200 {object} map[string]interface{} "JWT token successfully returned"
// @Failure 400 {object} models.ValidateErrorResponse "Bad Request: Validation error"
// @Failure 502 {object} models.ErrorResponse "Bad Gateway: Something bad happened"
// @Router /auth/login [post]
func SignInUser(c *fiber.Ctx) error {
	var input *models.SignInInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(input)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var user models.User
	result := database.DB.First(&user, "email = ?", strings.ToLower(input.Email))
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	config, _ := env.LoadConfig(".")

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(config.JwtSecret))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

// LogoutUser logs out a user.
// @Summary Log out a user
// @Description This endpoint allows a user to log out.
// @Tags Users
// @Produce json
// @Success 200 {object} map[string]interface{} "User successfully logged out"
// @Router /auth/logout [get]
func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

// GetMe retrieves the current logged in user.
// @Summary Get the current user
// @Description This endpoint returns the details of the current logged in user.
// @Tags Users
// @Produce json
// @Success 200 {object} map[string]interface{} "Current user data successfully returned"
// @Failure 401 {object} models.ErrorResponse "Unauthorized: User not logged in"
// @Router /users/me [get]
func GetMe(c *fiber.Ctx) error {
	user := c.Locals("user").(models.UserResponse)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": user}})
}
