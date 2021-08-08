package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/gofiber_svelte/model"
)

func FetchAll_mcompany(c *fiber.Ctx) error {
	result, err := model.FetchAll_mcompany()

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
