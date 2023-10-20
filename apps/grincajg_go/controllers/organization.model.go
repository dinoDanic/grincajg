package controllers

import (
	"grincajg/database"
	"grincajg/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	var input *models.CreateOrganizationInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(&input)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})
	}

	user := c.Locals("user").(models.UserResponse)

	newOrganization := models.Organization{
		UserID: &user.ID,
		Name:   input.Name,
	}

	result := database.DB.Create(&newOrganization)

	log.Println(result)

	return nil
}
