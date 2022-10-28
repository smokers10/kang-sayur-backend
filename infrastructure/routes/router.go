package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// test for connection
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"Message": "Server ON!", "App Version": "1.0 beta"})
	})

	// admin router
	AdminRouter(app)
}

func AdminRouter(app fiber.Router) {

}

func CustomerRouter(app fiber.Router) {

}
