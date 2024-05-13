package internal

import (
	"github.com/flambra/payment/internal/checkout"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).JSON(fiber.Map{
	// 		"project":     os.Getenv("PROJECT"),
	// 		"environment": os.Getenv("ENV"),
	// 		"version":     os.Getenv("BUILD_VERSION"),
	// 	})
	// })
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("../html/checkout.html")
	})
	app.Get("/success", func(c *fiber.Ctx) error {
		return c.SendFile("../html/success.html")
	})
	app.Get("/cancel", func(c *fiber.Ctx) error {
		return c.SendFile("../html/cancel.html")
	})
	app.Post("/checkout", checkout.Create)
}
