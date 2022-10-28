package main

import (
	"kang-sayur-backend/infrastructure/configuration"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// read config
	cnfg := configuration.ReadConfiguration()
	appConfig := cnfg.Application

	// init fiber framework
	app := fiber.New()

	// injector

	// init router

	// init server
	app.Listen(appConfig.Port)
}
