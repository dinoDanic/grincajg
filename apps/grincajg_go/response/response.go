package response

import (
	"grincajg/models"

	"github.com/gofiber/fiber/v2"
)

func ErrorResponse(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "fail",
		"message": message,
	})
}

func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   data,
	})
}

func ValidationErrorResponse(c *fiber.Ctx, errors []*models.ErrorResponse) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status": "fail",
		"errors": errors,
	})
}
