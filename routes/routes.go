package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Static("/about", "client/public", fiber.Static{
		Compress:  true,
		ByteRange: true,
		Browse:    true,
	})
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Render("index", fiber.Map{
	// 		"hello": "world",
	// 	})
	// })
	return app
}
