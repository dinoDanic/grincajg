package controllers

import (
	"grincajg/database"
	"grincajg/models"
	"grincajg/response"

	"github.com/gofiber/fiber/v2"
)

func CreateStore(c *fiber.Ctx) error {
	var input *models.CreateStoreInput

	if err := c.BodyParser(&input); err != nil {
		return response.ErrorResponse(c, err.Error())
	}

	err := models.ValidateStruct(input)

	if err != nil {
		return response.ErrorResponse(c, err.Error())
	}

	user := models.GetContextUser(c)
	database.DB.Preload("Organization").First(user)

	if user.Organization == nil {
		return response.ErrorResponse(c, "User does not belong to any organization")
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

func GetOrganizationStores(c *fiber.Ctx) error {
	user := models.GetContextUser(c)

	err := database.DB.Preload("Organization.Stores").Find(user).Error

	if err != nil {
		return response.ErrorResponse(c, err.Error())
	}

	if user.Organization == nil {
		return response.ErrorResponse(c, "User does not belong to any organization")
	}

	return response.SuccessResponse(c, models.FilterStoreRecords(user.Organization.Stores))

}
