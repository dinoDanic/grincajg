package controllers

import (
	// "grincajg/database"
	"grincajg/models"

	"github.com/gofiber/fiber/v2"
)

func CreateOrganization(c *fiber.Ctx) error {
	var input *models.CreateOrganizationInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	err := models.ValidateStruct(input)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": err})
	}

	user := c.Locals("user").(models.UserResponse)

	newOrganization := models.Organization{
		UserID: user.ID,
		Name:   input.Name,
	}

	orgEntry, error := newOrganization.SaveOrganization()

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"organization": orgEntry}})

}
