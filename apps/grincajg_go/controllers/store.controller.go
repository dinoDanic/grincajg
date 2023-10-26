package controllers

import (
	"grincajg/database"
	"grincajg/models"
	"grincajg/response"
	// "log"

	"github.com/gofiber/fiber/v2"
)

func CreateStore(c *fiber.Ctx) error {
	var input *models.CreateStoreInput

	if err := c.BodyParser(&input); err != nil {
		return response.ErrorResponse(c, err.Error())
	}

	err := models.ValidateStruct(input)

	if err != nil {
		return response.ValidationErrorResponse(c, err)
	}

	user := models.GetContextUser(c)
	database.DB.Preload("Organization").First(user)

	if user.Organization.ID == 0 {
		return response.ErrorResponse(c, "You dont belong to a organization")
	}

	newStore := models.Store{
		OrganizationID: user.Organization.ID,
		Name:           input.Name,
		Address:        input.Address,
	}

	storeEntry, error := newStore.SaveStore()

	if error != nil {
		return response.ErrorResponse(c, error.Error())
	}

	return response.SuccessResponse(c, models.FilterStoreRecord(*storeEntry))

}
