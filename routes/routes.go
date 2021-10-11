package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	app.Get("/knockknock", func(c *fiber.Ctx) error {
		return c.SendString("Who is there? 😾")
	})
}
