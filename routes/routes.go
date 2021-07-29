package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Static("/", "client/public")
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.JSON(fiber.Map{
	// 		"status":  fiber.StatusOK,
	// 		"message": "fiber svelte",
	// 	})
	// })
	return app
}
