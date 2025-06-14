package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/manish-id/manish-backend/controller"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to ManishOfficial")
	})

	// Product Routes
	ProductRoutes := app.Group("/products")
	ProductRoutes.Post("/create", controller.CreateProduct)
}
