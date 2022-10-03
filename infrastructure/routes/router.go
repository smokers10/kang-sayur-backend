package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// test for connection
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"Message": "Server ON!"})
	})
}
