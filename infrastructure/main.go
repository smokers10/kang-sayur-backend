package main

import (
	"kang-sayur-backend/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// init fiber framework
	app := fiber.New()

	// init router
	routes.Router(app)

	// init server
	app.Listen(":8000")
}
