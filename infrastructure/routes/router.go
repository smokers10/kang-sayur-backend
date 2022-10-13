package routes

import (
	"kang-sayur-backend/infrastructure/injector"
	"net/http"

	adminController "kang-sayur-backend/controller/admin"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App, injector injector.InjectorProvider) {
	// test for connection
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"Message": "Server ON!", "App Version": "1.0 beta"})
	})

	// api grouping
	api := app.Group("/api")

	// admin router
	AdminRouter(api, injector)

	// customer router
	CustomerRouter(app, injector)
}

func AdminRouter(app fiber.Router, injector injector.InjectorProvider) {
	// declarate neccesary data
	parent := app.Group("/admin")
	mc := adminController.MainController(injector)

	// authorization
	authController := mc.Authentication()
	authGroup := parent.Group("/auth")
	authGroup.Post("/request", authController.Request)
	authGroup.Post("/login", authController.Login)
}

func CustomerRouter(app fiber.Router, injector injector.InjectorProvider) {

}
