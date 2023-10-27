package controllers

import (
	"grincajg/database"
	"grincajg/models"
	"grincajg/response"
	"log"

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

	user := models.GetContextUser(c)

	newOrganization := models.Organization{
		AdminUserID: user.ID,
		Name:        input.Name,
	}

	orgEntry, error := newOrganization.SaveOrganization()

	if error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": error.Error()})
	}

	log.Println("orgEntryyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy")
	log.Println(orgEntry)

	user.OrganizationID = &orgEntry.ID

	result := database.DB.Save(&user)
	log.Println("resutlttttttttttttttttttttttttt")
	log.Println(result)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"organization": orgEntry}})

}

func GetOrganization(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	err := database.DB.Preload("Organization").Find(user).Error

	if err != nil {
		return response.ErrorResponse(c, err.Error())
	}

	if user.Organization == nil {
		return response.ErrorResponse(c, "You dont belong to an Organization")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"organization": models.FilterOrganizationRecord(*user.Organization)}})
}
