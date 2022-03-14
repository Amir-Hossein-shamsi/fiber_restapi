package routes

import (
	"github.com/gofiber/fiber/v2"
)

// type responseDefulat struct {
// 	status  int
// 	message string
// }

func Routes(app *fiber.App) {

	app.Get("/api/status", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return c.JSON("Ok Get")
	})
	app.Post("/api/book", func(c *fiber.Ctx) error {
		c.SendStatus(200)
		return c.JSON("Ok POST")
	})

}
