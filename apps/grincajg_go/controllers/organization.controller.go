package controllers

import (
	"grincajg/database"
	"grincajg/models"
	"grincajg/response"

	"github.com/gofiber/fiber/v2"
)

// CreateOrganization creates a new organization.
// @Summary Create a new organization
// @Description This endpoint creates a new organization and assigns the current logged in user as the admin of the organization.
// @Tags Organizations
// @Produce json
// @Param body body models.CreateOrganizationInput true "Organization creation input"
// @Success 201 {object} models.OrganizationRecord "Organization successfully created"
// @Failure 401 {object} models.Error "Unauthorized: User not logged in"
// @Failure 400 {object} models.ValidateErrorResponse "Bas request"
// @Router /organizations [post]
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

	user.OrganizationID = &orgEntry.ID

	result := database.DB.Save(&user)

	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": result.Error.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"organization": orgEntry}})

}

// GetOrganization retrieves the organization the current user belongs to.
// @Summary Get the current user's organization
// @Description This endpoint returns the details of the organization the current logged in user belongs to.
// @Tags Organizations
// @Produce json
// @Success 200 {object} models.Organization "Organization data successfully returned"
// @Failure 401 {object} models.Error "Unauthorized: User not logged in"
// @Failure 404 {object} models.Error "User does not belong to any organization"
// @Router /organization [get]
func GetOrganization(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)

	err := database.DB.Preload("Organization").Find(user).Error

	if err != nil {
		return response.ErrorResponse(c, err.Error())
	}

	if user.Organization == nil {
		return response.ErrorResponseNotFound(c, "User does not belong to any organization")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"organization": models.FilterOrganizationRecord(*user.Organization)}})
}
