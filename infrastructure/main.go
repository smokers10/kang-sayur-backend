package main

import (
	"kang-sayur-backend/infrastructure/configuration"
	"kang-sayur-backend/infrastructure/helper"
	"kang-sayur-backend/infrastructure/injector"
	"kang-sayur-backend/infrastructure/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// read config
	cnfg := configuration.ReadConfiguration()
	appConfig := cnfg.Application

	// init fiber framework
	app := fiber.New()

	// call injector
	injector := injector.Injector()

	helper.MongoBuilder(injector.Database).CollectionBuilder()

	// init router
	routes.Router(app, *injector)

	// init server
	app.Listen(appConfig.Port)
}
