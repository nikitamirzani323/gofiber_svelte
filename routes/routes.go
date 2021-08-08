package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nikitamirzani323/gofiber_svelte/controller"
)

func Init() *fiber.App {
	app := fiber.New(fiber.Config{
		Concurrency: 256 * 1024,
	})

	app.Use(logger.New())
	app.Static("/", "client/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})

	app.Get("/api/company", controller.FetchAll_mcompany)
	return app
}
