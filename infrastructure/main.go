package main

import (
	"kang-sayur-backend/infrastructure/configuration"
	"kang-sayur-backend/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// read config
	cnfg := configuration.ReadConfiguration()
	appConfig := cnfg.Application

	// init fiber framework
	app := fiber.New()

	// init router
	routes.Router(app)

	// init server
	app.Listen(appConfig.Port)
}
