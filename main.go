package main

import (
	"github.com/manish-id/manish-backend/config"
	"github.com/manish-id/manish-backend/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Koneksi ke MongoDB
	config.ConnectMongo()

	// Setup semua route (endpoint API)
	routes.SetupRoutes(app)

	// Jalankan server di port 3000
	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
